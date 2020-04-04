package persistance

import (
	"bettertomorrow/context/company/domain"
	"bettertomorrow/common/db"
	"sync"
)

type CompanyRepository interface {
	Insert(company *domain.Company) error
}

type CompanyRepositoryImpl struct {
	db *db.DbConnection
}

/* -- singleton for DI -- */

var companyRepositoryInstance *CompanyRepositoryImpl
var once sync.Once

func ProvideCompanyRepositoryImpl() *CompanyRepositoryImpl {
	once.Do(func() {
		dbConnection := db.ProvideDbConnection()
		companyRepositoryInstance = &CompanyRepositoryImpl{dbConnection }
	})
	return companyRepositoryInstance
}

/* ---- */

func (crImpl *CompanyRepositoryImpl) Insert(company *domain.Company) error {
	return nil
}


