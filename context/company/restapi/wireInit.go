// +build wireinject

package restapi

import (
	"bettertomorrow/context/company/application"
	"github.com/google/wire"
)

func NewCompanyRouter() (*CompanyRouter, error) {
	wire.Build(
		application.ProvideCompanyServiceImpl, 
		application.ProvideCompanyProductsServiceImpl, 
		instantiateCompanyRouter,
	)
	return &CompanyRouter{}, nil
}

func NewProductRouter() (*ProductRouter, error) {
	wire.Build(
		application.ProvideProductServiceImpl,
		wire.Bind(new(application.ProductService), new(*application.ProductServiceImpl)),
		instantiateProductRouter,
	)
	return &ProductRouter{}, nil
}