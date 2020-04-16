package restapi

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	// "github.com/pkg/errors"

	//"time"
	"bettertomorrow/context/customer/application"
	"bettertomorrow/context/customer/domain"
)

type CustomerRouter struct {
	customerService                   *application.CustomerServiceImpl
	customerServiceWithWalletsService *application.CustomerWalletsServiceImpl
}

//TODO: pass interface
func instantiateCustomerRouter(customerService *application.CustomerServiceImpl, customerServiceWithWalletsService *application.CustomerWalletsServiceImpl) *CustomerRouter {
	return &CustomerRouter{customerService, customerServiceWithWalletsService}
}

func (cr *CustomerRouter) AddRoutes(apiRoutes *echo.Group) {
	apiRoutes.GET("/customers", func(c echo.Context) error {
		customers, err := cr.customerService.FindAllCompanies()
		if err != nil {
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
			return err
		}

		return c.JSON(http.StatusOK, "OK")
	})

	apiRoutes.PUT("/customers", func(c echo.Context) error {
		updatedCustomer := &domain.Customer{}
		err := c.Bind(updatedCustomer)
		if err != nil {
			return err
		}
		fmt.Printf("update customer: %v \n", updatedCustomer)
		
		err = cr.customerService.UpdateCustomer(updatedCustomer)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, "OK")
	})

	apiRoutes.DELETE("/customers/:id", func(c echo.Context) error {
		idStr := c.Param("id")
		id, err := strconv.ParseUint(idStr, 10, 32)
		if err != nil {
			return err
		}
		fmt.Printf("delete customer with id: %v \n", id)

		err = cr.customerService.DeleteCustomer(uint(id))
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, "OK")
	})
}
