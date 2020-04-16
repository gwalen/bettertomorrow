/**
 * Arificialy created struct to have dependecies disturbutted many structs (more than 1:1 (ex.: 1 repository <-> 1 service))
 */

package application

import (
	"bettertomorrow/context/company/domain"
	"bettertomorrow/context/company/persistance"
	"sync"
)

 type CompanyProductsServiceImpl struct {
	 companyRepository persistance.CompanyRepository
	 productRepository persistance.ProductRepository
 }

 /* -- singleton for DI -- */

var companyProductsServiceInstance *CompanyProductsServiceImpl
var onceForCompanyProductsService sync.Once

func ProvideCompanyProductsServiceImpl() *CompanyProductsServiceImpl {
	onceForCompanyProductsService.Do(func() {
		companyProductsServiceInstance = &CompanyProductsServiceImpl{
			companyRepository: persistance.ProvideCompanyRepositoryImplGorm(),
			productRepository: persistance.ProvideProductRepositoryImpl(),
		}
	})

	return companyProductsServiceInstance
} 

 /* ---- */

 func (cpsImpl *CompanyProductsServiceImpl) FindCompanyWithProducts() ([]domain.Company, error) {
	 return cpsImpl.companyRepository.FindWithProducts()
 }

 