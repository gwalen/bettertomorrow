/**
 * Arificialy created struct to have dependecies disturbutted many structs (more than 1:1 (ex.: 1 repository <-> 1 service))
 */

package application

import (
	"bettertomorrow/context/customer/domain"
	"bettertomorrow/context/customer/persistance"
	"sync"
)

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

 