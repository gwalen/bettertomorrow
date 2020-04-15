package application

import (
	"bettertomorrow/context/company/domain"
	"bettertomorrow/context/company/persistance"
	"sync"
)

type ProductService interface {
	CreateProduct(company *domain.Product) error
	UpdateProduct(company *domain.Product) error
	DeleteProduct(id uint) error
	FindAllProducts() ([]domain.Product, error)
}

type ProductServiceImpl struct {
	productRepository persistance.ProductRepository
}

/* -- singleton for DI -- */

var productServiceInstance *ProductServiceImpl
var onceForProductService sync.Once

func ProvideProductServiceImpl() *ProductServiceImpl {
	onceForProductService.Do(func() {
		var productRepository persistance.ProductRepository
		productRepositoryImpl := persistance.ProvideProductRepositoryImpl()
		productRepository = productRepositoryImpl
		productServiceInstance = &ProductServiceImpl{productRepository}
	})
	return productServiceInstance
}

/* ---- */

func (csImpl *ProductServiceImpl) CreateProduct(company *domain.Product) error {
	return csImpl.productRepository.Insert(company)
}

func (csImpl *ProductServiceImpl) UpdateProduct(company *domain.Product) error {
	return csImpl.productRepository.InsertOrUpdate(company)
}

func (csImpl *ProductServiceImpl) DeleteProduct(id uint) error {
	return csImpl.productRepository.Delete(id)
}

func (csImpl *ProductServiceImpl) FindAllProducts() ([]domain.Product, error) {
	products, err := csImpl.productRepository.FindAll()
	return products, err
}
