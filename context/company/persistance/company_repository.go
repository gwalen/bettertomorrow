package persistance

import (
	"bettertomorrow/context/company/domain"
)

type CompanyRepository interface {
	Insert(company *domain.Company) error
	InsertOrUpdate(company *domain.Company) error
	Delete(id uint) error
	FindAll() ([]domain.Company, error)
	FindWithProducts(companyName string) ([]domain.CompanyWithProducts, error)
}
