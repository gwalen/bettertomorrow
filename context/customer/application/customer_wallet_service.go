/**
 * Arificialy created struct to have dependecies disturbutted among many structs (more than 1:1 (ex.: 1 repository <-> 1 service))
 */

package application

import (
	"bettertomorrow/common/logger"
	. "bettertomorrow/common/util/concurrency"
	"bettertomorrow/context/customer/domain"
	"bettertomorrow/context/customer/persistance"

	"errors"
	// "fmt"
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

	SafeGoWithErrorChan(logCws, func(){ ReadInput(input, readerFunc) }, errorChan)
	SafeGoWithErrorChan(logCws, func(){ StartWorkerPool(logCws, workerNumber, input, output, errorChan, workerFunc) }, errorChan)
	SafeGo(logCws, func(){ HandleResourcesOnError(logCws, input, output, errorChan) })
	WriteOutput(output, errorChan, resultWriterFunc)

	return results, nil
} 
