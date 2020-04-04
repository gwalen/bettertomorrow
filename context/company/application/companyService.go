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
	companyRepository *persistance.CompanyRepositoryImpl //TODO test with interface (than you have to omit *)
}

/* -- singleton for DI -- */

var companyServiceInstance *CompanyServiceImpl
var once sync.Once

func ProvideCompanyServiceImpl() *CompanyServiceImpl {
	once.Do(func() {
		companyRepository := persistance.ProvideCompanyRepositoryImpl()
		companyServiceInstance = &CompanyServiceImpl{companyRepository}
	})
	return companyServiceInstance
}

/* ---- */

func (csImpl *CompanyServiceImpl) CreateCompany(company *domain.Company) error {
	// TODO
	return csImpl.companyRepository.Insert(company)
}

func (csImpl *CompanyServiceImpl) FindAllCompanies() ([]domain.Company, error) {
	companiesMock := []domain.Company{
		{domain.Address{"street-1", "12", "02-744", "wawa", "PL",},100, "test-1", "tax-id-test"},
		{domain.Address{}, 101, "test-2", "tax-id-test"},
	}
	return companiesMock, nil
}
