package persistance

import (
	"bettertomorrow/common/dbgorm"
	"bettertomorrow/context/company/domain"
	"sync"

	"github.com/jinzhu/gorm"
)

type CompanyRepository interface {
	Insert(company *domain.Company) error
	InsertOrUpdate(company *domain.Company) error
	Delete(id uint) error
	FindAll() ([]domain.Company, error)
	FindWithProducts() ([]domain.Company, error)
	FindWithProductsRawSql() ([]domain.CompanyWithProducts, error)
}

type CompanyRepositoryImplGorm struct {
	db *gorm.DB
}

/* -- singleton for DI -- */

var companyRepositoryInstanceGorm *CompanyRepositoryImplGorm
var onceForCompanyRepositoryGorm sync.Once

func ProvideCompanyRepositoryImplGorm() *CompanyRepositoryImplGorm {
	onceForCompanyRepositoryGorm.Do(func() {
		dbHandle := dbgorm.DB()
		companyRepositoryInstanceGorm = &CompanyRepositoryImplGorm{dbHandle}
	})
	return companyRepositoryInstanceGorm
}

/* ---- */

func (crImpl *CompanyRepositoryImplGorm) Insert(company *domain.Company) error {
	company.ID = 0 // setting to 0 will trigger auto increment
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
	err := crImpl.db.Find(&companies).Error // Products field in Company struct will be nil becuse we only pull companies without associations
	return companies, err
}

func (crImpl *CompanyRepositoryImplGorm) FindWithProducts() ([]domain.Company, error) {
	var companies []domain.Company
	err := crImpl.db.Preload("Products").Table("companies").Joins("join products on products.company_id = companies.id").Find(&companies).Error

	return companies, err
}

func (crImpl *CompanyRepositoryImplGorm) FindWithProductsRawSql() ([]domain.CompanyWithProducts, error) {
	var companies []domain.Company
	err := crImpl.db.Preload("Products").Table("companies").Joins("join products on products.company_id = companies.id").Find(&companies).Error

	var companiesWithProducts []domain.CompanyWithProducts
	for _, elem := range companies {
		products := elem.Products
		elem.Products = nil // as we return products in CompanyWithProducts object not in Company
		cp := domain.CompanyWithProducts{elem, products}
		companiesWithProducts = append(companiesWithProducts, cp)
	}

	return companiesWithProducts, err
}
