// +build wireinject

package restapi

import (
	"bettertomorrow/context/customer/application"
	"github.com/google/wire"
)

func NewCustomerRouter() *CustomerRouter {
	wire.Build(
		application.ProvideCustomerServiceImpl,
		application.ProvideCustomerWalletsServiceImpl,
		instantiateCustomerRouter,
	)
	return &CustomerRouter{}
}

func NewWalletRouter() *WalletRouter {
	//TODO: any way to handle errors ?
	wire.Build(
		application.ProvideWalletServiceImpl,
		wire.Bind(new(application.WalletService), new(*application.WalletServiceImpl)),
		instantiateWalletRouter,
	)
	return &WalletRouter{}
}
