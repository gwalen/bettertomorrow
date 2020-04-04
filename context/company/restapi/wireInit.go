//+build wireinject

package restapi

import (
	"bettertomorrow/context/company/application"
	"github.com/google/wire"
)

func NewCompanyRouter() (*CompanyRouter, error) {
	wire.Build(application.ProvideCompanyServiceImpl, instantiateCompanyRouter)
	return &CompanyRouter{}, nil
}