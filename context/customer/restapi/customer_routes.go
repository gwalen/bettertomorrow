package restapi

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"bettertomorrow/common/logger"

	"bettertomorrow/context/customer/application"
	"bettertomorrow/context/customer/domain"
)

var logCustomer = logger.ProvideLogger()

type CustomerRouter struct {
	customerService                   *application.CustomerServiceImpl
	customerServiceWithWalletsService *application.CustomerWalletsServiceImpl
}


//TODO: add logger fgacade so I can chnage ubderlying implementation (zap, or zerologger)
//TODO: pass interface
func instantiateCustomerRouter(customerService *application.CustomerServiceImpl, customerServiceWithWalletsService *application.CustomerWalletsServiceImpl) *CustomerRouter {
	return &CustomerRouter{customerService, customerServiceWithWalletsService}
}

func (cr *CustomerRouter) AddRoutes(apiRoutes *echo.Group) {
	apiRoutes.GET("/customers", func(c echo.Context) error {
		logCustomer.Info("Find all customers")
		customers, err := cr.customerService.FindAllCustomers()
		if err != nil {
			logCustomer.Error("Failed to fetch customers", err)
			return err
		}
		return c.JSON(http.StatusOK, customers)
	})

	apiRoutes.GET("/customers/wallets", func(c echo.Context) error {
		customersWithProdcut, err := cr.customerServiceWithWalletsService.FindCustomerWithWallets()
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, customersWithProdcut)
	})

	apiRoutes.POST("/customers", func(c echo.Context) error {
		newCustomer := &domain.Customer{}
		err := c.Bind(newCustomer)
		if err != nil {
			return err
		}

		err = cr.customerService.CreateCustomer(newCustomer)
		if err != nil {
			logCustomer.Error(fmt.Sprintf("Failed to add customer %v", newCustomer), err)
			return err
		}
		logCustomer.Info(fmt.Sprintf("New customer %v\n", newCustomer))

		return c.JSON(http.StatusOK, "OK")
	})

	apiRoutes.PUT("/customers", func(c echo.Context) error {
		updatedCustomer := &domain.Customer{}
		err := c.Bind(updatedCustomer)
		if err != nil {
			return err
		}

		err = cr.customerService.UpdateCustomer(updatedCustomer)
		if err != nil {
			logCustomer.Error(fmt.Sprintf("Failed to update customer %v", updatedCustomer), err)
			return err
		}

		return c.JSON(http.StatusOK, "OK")
	})

	apiRoutes.DELETE("/customers/:id", func(c echo.Context) error {
		idStr := c.Param("id")
		id, err := strconv.ParseUint(idStr, 10, 32)
		if err != nil {
			logCustomer.Error(fmt.Sprintf("Failed to delete customer with id = %v", id), err)
			return err
		}
		logCustomer.Info(fmt.Sprintf("delete customer with id: %v\n", id))

		err = cr.customerService.DeleteCustomer(uint(id))
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, "OK")
	})
}
