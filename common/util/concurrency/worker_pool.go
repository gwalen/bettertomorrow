package concurrency

import (
	"bettertomorrow/common/logger"
	"fmt"
	"sync"
)

func Worker(wg *sync.WaitGroup, intput chan interface{}, output chan interface{}, workerFunc func(interface{}) interface{}) {
	for value := range intput {
		result := workerFunc(value)
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
	for _, inputValue := range inputData {
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
		if executionResult != "done" { // error in sub routine, close resuources
			SafeClose(log, input)	
			SafeClose(log, output)
			log.Warn(fmt.Sprintf("Error in subroutine, release resources: closing channels; error: %v", executionResult))
		} else {
			log.Info("All jobs finsed")
		}
	}
}