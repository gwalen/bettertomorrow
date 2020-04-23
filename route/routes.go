package route

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"time"
	restapiCompany  "bettertomorrow/context/company/restapi"
	restapiCustomer "bettertomorrow/context/customer/restapi"
	restapiEmployee "bettertomorrow/context/employee/restapi"
)

func Init() *echo.Echo {

	echoServer := echo.New()

	echoServer.Use(middleware.Logger())

	apiRoutes := echoServer.Group("/api")

	addUtilRoutes(apiRoutes)
	addCompanyRoutes(apiRoutes)
	addCustomerRoutes(apiRoutes)
	addEmployeeRoutes(apiRoutes)

	return echoServer
}

func addUtilRoutes(apiRoutes *echo.Group) {
	apiRoutes.Group("/health").GET("/check", func(c echo.Context) error {
		return c.String(http.StatusOK, time.Now().String())
	})
}

func addCompanyRoutes(apiRoutes *echo.Group) {
	router, _ := restapiCompany.NewCompanyRouter()
	router.AddRoutes(apiRoutes)
	addProductRoutes(apiRoutes)
}

func addProductRoutes(apiRoutes *echo.Group) {
	router, _ := restapiCompany.NewProductRouter()
	router.AddRoutes(apiRoutes)
}

func addCustomerRoutes(apiRoutes *echo.Group) {
	router := restapiCustomer.NewCustomerRouter()
	router.AddRoutes(apiRoutes)
	addWalletRoutes(apiRoutes)
}

func addWalletRoutes(apiRoutes *echo.Group) {
	router := restapiCustomer.NewWalletRouter()
	router.AddRoutes(apiRoutes)
}

func addRoleRoutes(apiRoutes *echo.Group) {
	router := restapiEmployee.NewRoleRouter()
	router.AddRoutes(apiRoutes)
}

func addEmployeeRoutes(apiRoutes *echo.Group) {
	router := restapiEmployee.NewEmployeeRouter()
	router.AddRoutes(apiRoutes)
	addRoleRoutes(apiRoutes)
}


