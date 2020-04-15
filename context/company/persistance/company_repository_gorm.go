package persistance

import (
	"bettertomorrow/common/dbgorm"
	"bettertomorrow/context/company/domain"
	"fmt"
	"sync"

	"github.com/jinzhu/gorm"
)

type CompanyRepositoryImplGorm struct {
	db *gorm.DB
}

/* -- singleton for DI -- */

var companyRepositoryInstanceGorm *CompanyRepositoryImplGorm
var onceForCompanyRepositoryGorm sync.Once

func ProvideCompanyRepositoryImplGorm() *CompanyRepositoryImplGorm {
	onceForCompanyRepositoryGorm.Do(func() {
		dbConnection := dbgorm.DB()
		companyRepositoryInstanceGorm = &CompanyRepositoryImplGorm{dbConnection}
	})
	fmt.Printf("INIT GORM REPO %v\n", companyRepositoryInstanceGorm)
	return companyRepositoryInstanceGorm
}

/* ---- */

func (crImpl *CompanyRepositoryImplGorm) Insert(company *domain.Company) error {
	company.ID = 0 // setting to 0 will trigger auto increment
	// return db.DB().Create(company).Error //TODO: refactor and test
	return crImpl.db.Create(company).Error
}

func (crImpl *CompanyRepositoryImplGorm) InsertOrUpdate(company *domain.Company) error {
	return crImpl.db.Save(company).Error
}

/** 
 * gorm has also soft delete by default (when usingg DeletedAt field)
 */
func (crImpl *CompanyRepositoryImplGorm) Delete(id uint) error {
	//passing some value to Delete(..) is mandatory even if not used
	return crImpl.db.Where("id = ?", id).Delete(domain.Company{}).Error
}

func (crImpl *CompanyRepositoryImplGorm) FindAll() ([]domain.Company, error) {
	var companies []domain.Company
	// TODO:
	// err := crImpl.db.Find(companies).Error //interstingly error is reported in seversl different go routines althought there was just one call ?
	err := crImpl.db.Find(&companies).Error
	return companies, err
}

func (crImpl *CompanyRepositoryImplGorm) FindWithProducts(companyName string) ([]domain.CompanyWithProducts, error) {
	var companies []domain.Company
	err := crImpl.db.Preload("Products").Table("companies").Joins("join products on products.company_id = companies.id").Find(&companies).Error
		// Where("companies.name like ?", "'" + companyName + "%'").
		// Where("companies.name = ?", companyName).
		// Where("companies.id = ?", 1).
		

	var companiesWithProducts []domain.CompanyWithProducts
	for _, elem := range companies {
		cp := domain.CompanyWithProducts{elem, elem.Products}
		companiesWithProducts = append(companiesWithProducts, cp)
	}

	return companiesWithProducts, err
}