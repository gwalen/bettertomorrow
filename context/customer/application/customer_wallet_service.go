/**
 * Arificialy created struct to have dependecies disturbutted among many structs (more than 1:1 (ex.: 1 repository <-> 1 service))
 */

package application

import (
	"bettertomorrow/common/logger"
	. "bettertomorrow/common/util/concurrency"
	"bettertomorrow/context/customer/domain"
	"bettertomorrow/context/customer/persistance"

	// "errors"
	"fmt"
	"sync"
)

var logCws = logger.ProvideLogger()

//TODO: add interface ?
type CustomerWalletsServiceImpl struct {
	 customerRepository persistance.CustomerRepository
	 walletRepository persistance.WalletRepository
 }

 /* -- singleton for DI -- */

var customerWalletsServiceInstance *CustomerWalletsServiceImpl
var onceForCustomerWalletsService sync.Once

func ProvideCustomerWalletsServiceImpl() *CustomerWalletsServiceImpl {
	onceForCustomerWalletsService.Do(func() {
		customerWalletsServiceInstance = &CustomerWalletsServiceImpl{
			customerRepository: persistance.ProvideCustomerRepositoryImpl(),
			walletRepository: persistance.ProvideWalletRepositoryImpl(),
		}
	})

	return customerWalletsServiceInstance
} 

 /* ---- */

 func (cpsImpl *CustomerWalletsServiceImpl) FindCustomerWithWallets() ([]domain.CustomerWithWallets, error) {
	 return cpsImpl.customerRepository.FindWithWallets()
 }


var workerNumber = 4

func aggregateSavings(customerWallets domain.CustomerWithWallets) domain.AggregatedWallet {
	var totoalUnits float64 
	var currencies []string	
	for _, wallet := range customerWallets.Wallet {
		totoalUnits += float64(wallet.Amount)
		currencies = append(currencies, wallet.Currency)
	}
	return domain.AggregatedWallet{currencies, totoalUnits}
}

func worker(wg *sync.WaitGroup, intput chan interface{}, output chan interface{}, workerFunc func(interface{}) interface{}) {
	fmt.Printf("worker: waiting for value \n")
	for value := range intput {
		fmt.Printf("worker value from input : %v \n", value)
		result := workerFunc(value)
		fmt.Printf("worker result: %v, for input : %v \n", result, value)
		output <- result
	} 
	wg.Done()
}

func startWorkerPool(poolSize int, input chan interface{}, output chan interface{}, errorChan chan string, workerFunc func(interface{}) interface{}) {
	var wg sync.WaitGroup
	for i := 0; i < poolSize; i++ {
		wg.Add(1)
		SafeGoWithErrorChan(logCws, func(){ worker(&wg, input, output, workerFunc) }, errorChan)	
	}
	wg.Wait() 
	SafeClose(logCws, output)
}

func readInput(input chan interface{}, readerFunc func() []interface{}) {
	inputData := readerFunc()
	fmt.Printf("read input  data slice: : %v \n", inputData)
	for _, inputValue := range inputData {
		fmt.Printf("read input data iteration : %v \n", inputValue)
		input <- inputValue
	}
	close(input)
}

func writeOutput(output chan interface{}, errorChan chan string, resultWriterFunc func(interface{})) {
	for resultData := range output {
		resultWriterFunc(resultData)
	}
	errorChan <- "done"
	close(errorChan)	
}

func handleResourcesOnError(input chan interface{}, output chan interface{}, errorChan chan string) {
	for executionResult := range errorChan {
		if executionResult != "done" { // error in sub routine, close resuources and panic
			SafeClose(logCws, input)	
			SafeClose(logCws, output)
			logCws.Warn(fmt.Sprintf("Error in subroutine - closing channels, error: %v", executionResult))
		} else {
			logCws.Info("All jobs finsed")
		}
	}
}

func (cpsImpl *CustomerWalletsServiceImpl) AggregateCustomerSavings() ([]domain.AggregatedWallet, error) {
	input := make(chan interface{}, workerNumber) 
	output := make(chan interface{}, workerNumber) 
	errorChan := make(chan string)
	var results []domain.AggregatedWallet

	readerFunc := func() []interface{} { 
		customersWallets, err := cpsImpl.customerRepository.FindWithWallets() 
		if err != nil {
			logCws.Error("failed to read customer wallets", err)
			return nil			
		}

		data := make([]interface{}, len(customersWallets))
		for i, cs := range customersWallets { data[i] = cs }
		// panic(errors.New("TESTING PANIC IN SUB ROUTINE"))

		return data
	} 

	workerFunc := func(value interface{}) interface{} { 
		return aggregateSavings(value.(domain.CustomerWithWallets)) 
	}

	resultWriterFunc := func(resultData interface{}) {
		results = append(results, resultData.(domain.AggregatedWallet))
	}	

	SafeGoWithErrorChan(logCws, func(){ readInput(input, readerFunc) }, errorChan)
	SafeGoWithErrorChan(logCws, func(){ startWorkerPool(workerNumber, input, output, errorChan, workerFunc) }, errorChan)
	SafeGo(logCws, func(){ handleResourcesOnError(input, output, errorChan) })
	writeOutput(output, errorChan, resultWriterFunc)

	return results, nil
} 

 