package application

import (
	"bettertomorrow/context/company/domain"
	"bettertomorrow/context/company/persistance"
	"sync"
)

type CompanyService interface {
	CreateCompany(company *domain.Company) error
}

type CompanyServiceImpl struct {
	companyRepository persistance.CompanyRepository
}

/* -- singleton for DI -- */

var companyServiceInstance *CompanyServiceImpl
var onceForCompantService sync.Once

func ProvideCompanyServiceImpl() *CompanyServiceImpl {
	onceForCompantService.Do(func() {
		// companyRepository := persistance.ProvideCompanyRepositoryImplGorm()
		companyRepository := persistance.ProvideCompanyRepositoryImplXorm()
		companyServiceInstance = &CompanyServiceImpl{companyRepository}
	})
	return companyServiceInstance
}

/* ---- */

func (csImpl *CompanyServiceImpl) CreateCompany(company *domain.Company) error {
	return csImpl.companyRepository.Insert(company)
}

func (csImpl *CompanyServiceImpl) UpdateCompany(company *domain.Company) error {
	return csImpl.companyRepository.InsertOrUpdate(company)
}

func (csImpl *CompanyServiceImpl) DeleteCompany(id uint) error {
	return csImpl.companyRepository.Delete(id)
}

func (csImpl *CompanyServiceImpl) FindAllCompanies() ([]domain.Company, error) {
	companies, err := csImpl.companyRepository.FindAll()
	return companies, err
}
