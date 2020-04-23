package restapi

import (
	"net/http"

	"bettertomorrow/context/customer/domain"
	"bettertomorrow/context/customer/application"

	"github.com/labstack/echo/v4"
)

type WalletRouter struct {
	walletService application.WalletService
}

func instantiateWalletRouter(walletService application.WalletService) *WalletRouter {
	return &WalletRouter{walletService}
}

func (wr *WalletRouter) AddRoutes(apiRoutes *echo.Group) {
	apiRoutes.GET("/wallets", func(c echo.Context) error {
		wallets, err := wr.walletService.FindAllWallets()
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, wallets)
	})

	apiRoutes.POST("/wallets", func(c echo.Context) error {
		var newWallet *domain.Wallet
		if err := c.Bind(newWallet); err != nil {
			return err
		} 
		wr.walletService.CreateWallet(newWallet)	
		return c.JSON(http.StatusOK, "OK")		
	})
}
