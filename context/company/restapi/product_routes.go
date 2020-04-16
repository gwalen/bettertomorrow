package restapi

import (
	"net/http"

	"bettertomorrow/context/company/application"

	"github.com/labstack/echo/v4"
)

type ProductRouter struct {
	productService application.ProductService
}

func instantiateProductRouter(productService application.ProductService) *ProductRouter {
	return &ProductRouter{productService}
}

func (pr *ProductRouter) AddRoutes(apiRoutes *echo.Group) {
	apiRoutes.GET("/products", func(c echo.Context) error {
		products, err := pr.productService.FindAllProducts()
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, products)
	})
}