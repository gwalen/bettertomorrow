package persistance

import (
	"bettertomorrow/common/dbgorm"
	"bettertomorrow/context/company/domain"
	"sync"

	"github.com/jinzhu/gorm"
)

type ProductRepository interface {
	Insert(product *domain.Product) error
	InsertOrUpdate(product *domain.Product) error
	Delete(id uint) error
	FindAll() ([]domain.Product, error)
}

type ProductRepositoryImpl struct {
	db *gorm.DB
}

/* -- singleton for DI -- */

var productRepositoryInstance *ProductRepositoryImpl
var onceForProductRepository sync.Once

func ProvideProductRepositoryImpl() *ProductRepositoryImpl {
	onceForProductRepository.Do(func() {
		dbHandle := dbgorm.DB()
		productRepositoryInstance = &ProductRepositoryImpl{dbHandle}
	})
	return productRepositoryInstance
}

/* ---- */

func (prImpl *ProductRepositoryImpl) Insert(product *domain.Product) error {
	return prImpl.db.Create(product).Error
}

func (prImpl *ProductRepositoryImpl) InsertOrUpdate(product *domain.Product) error {
	return prImpl.db.Save(product).Error
}

func (prImpl *ProductRepositoryImpl) Delete(id uint) error {
	return prImpl.db.Where("id = ?", id).Delete(domain.Product{}).Error
}

func (prImpl *ProductRepositoryImpl) FindAll() ([]domain.Product, error) {
	var products []domain.Product
	error := prImpl.db.Find(&products).Error
	return products, error
}
