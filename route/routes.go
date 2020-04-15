package route

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"time"
	"bettertomorrow/context/company/restapi"
)

func Init() *echo.Echo {

	echoServer := echo.New()

	echoServer.Use(middleware.Logger())

	apiRoutes := echoServer.Group("/api")

	addUtilRoutes(apiRoutes)
	addCompanyRoutes(apiRoutes)
	addProductRoutes(apiRoutes)

	return echoServer
}

func addUtilRoutes(apiRoutes *echo.Group) {
	apiRoutes.Group("/health").GET("/check", func(c echo.Context) error {
		return c.String(http.StatusOK, time.Now().String())
	})
}

func addCompanyRoutes(apiRoutes *echo.Group) {
	router, _ := restapi.NewCompanyRouter()
	router.AddRoutes(apiRoutes)
}

func addProductRoutes(apiRoutes *echo.Group) {
	router, _ := restapi.NewProductRouter()
	router.AddRoutes(apiRoutes)
}


