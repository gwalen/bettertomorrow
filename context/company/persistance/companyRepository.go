package persistance

import (
	"bettertomorrow/context/company/domain"
	"bettertomorrow/common/db"
	"sync"

	"github.com/jinzhu/gorm"
)

type CompanyRepository interface {
	Insert(company *domain.Company) error
}

type CompanyRepositoryImpl struct {
	db *gorm.DB
}

/* -- singleton for DI -- */

var companyRepositoryInstance *CompanyRepositoryImpl
var once sync.Once

func ProvideCompanyRepositoryImpl() *CompanyRepositoryImpl {
	once.Do(func() {
		dbConnection := db.DB()
		companyRepositoryInstance = &CompanyRepositoryImpl{dbConnection }
	})
	return companyRepositoryInstance
}

/* ---- */

func (crImpl *CompanyRepositoryImpl) Insert(company *domain.Company) error {
	return db.DB().Create(company).Error
}

func (crImpl *CompanyRepositoryImpl) FindAll() ([]domain.Company, error) {
	var companies []domain.Company
	// TODO:
	// err := db.DB().Find(companies).Error //interstingly error is reported in seversl different go routines althought there was just one call ?
	err := db.DB().Find(&companies).Error  
	return companies, err
}
