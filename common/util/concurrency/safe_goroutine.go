package concurrency

import (
	"bettertomorrow/common/logger"
	"fmt"
)

/**
 * Usage: 
 * safeGo(logger, func() {
 *		functionToRunInGoRoutine(...)
 * }) 
 *
 */

func SafeGo(log logger.Logger, f func()) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				err := r.(error)
				log.Error("parallel go routnie panic", err)
			}
		}()

		f()
	}()
}

func SafeGoWithErrorChan(log logger.Logger, f func(), errorChan chan string) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				err := r.(error)
				log.Error("parallel go routnie panic", err)
				errorChan <- err.Error()
			}
		}()

		f()
	}()
}

func SafeClose(log logger.Logger, c chan interface{}) {
	defer func() {
		if r := recover(); r != nil {
			log.Warn(fmt.Sprintf("Error closing channel error: %v", r))
		}
	}()
	close(c)
}