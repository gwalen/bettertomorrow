package concurrency

import (
	"bettertomorrow/common/logger"
	"fmt"
	"sync"
)

func Worker(wg *sync.WaitGroup, intput chan interface{}, output chan interface{}, workerFunc func(interface{}) interface{}) {
	fmt.Printf("worker: waiting for value \n")
	for value := range intput {
		fmt.Printf("worker value from input : %v \n", value)
		result := workerFunc(value)
		fmt.Printf("worker result: %v, for input : %v \n", result, value)
		output <- result
	} 
	wg.Done()
}

func StartWorkerPool(
	log logger.Logger,
	poolSize int,
	input chan interface{},
	output chan interface{},
	errorChan chan string,
	workerFunc func(interface{}) interface{},
){
	var wg sync.WaitGroup
	for i := 0; i < poolSize; i++ {
		wg.Add(1)
		SafeGoWithErrorChan(log, func(){ Worker(&wg, input, output, workerFunc) }, errorChan)	
	}
	wg.Wait() 
	SafeClose(log, output)
}

func ReadInput(input chan interface{}, readerFunc func() []interface{}) {
	inputData := readerFunc()
	fmt.Printf("read input  data slice: : %v \n", inputData)
	for _, inputValue := range inputData {
		fmt.Printf("read input data iteration : %v \n", inputValue)
		input <- inputValue
	}
	close(input)
}

func WriteOutput(output chan interface{}, errorChan chan string, resultWriterFunc func(interface{})) {
	for resultData := range output {
		resultWriterFunc(resultData)
	}
	errorChan <- "done"
	close(errorChan)	
}

func HandleResourcesOnError(log logger.Logger, input chan interface{}, output chan interface{}, errorChan chan string) {
	for executionResult := range errorChan {
		if executionResult != "done" { // error in sub routine, close resuources and panic
			SafeClose(log, input)	
			SafeClose(log, output)
			log.Warn(fmt.Sprintf("Error in subroutine, release resources: closing channels; error: %v", executionResult))
		} else {
			log.Info("All jobs finsed")
		}
	}
}