package restapi

import (
	// "fmt"
	"fmt"
	"net/http"
	// "strconv"

	"bettertomorrow/context/company/application"

	"github.com/labstack/echo/v4"
	// "bettertomorrow/context/company/domain"
)

//TODO: do like this (type shuould have an interface not a concreate implementation)
type ProductRouter struct {
	productService application.ProductService
}

func instantiateProductRouter(productService application.ProductService) *ProductRouter {
	return &ProductRouter{productService}
}

func (pr *ProductRouter) AddRoutes(apiRoutes *echo.Group) {
	apiRoutes.GET("/products", func(c echo.Context) error {
		products, err := pr.productService.FindAllProducts()
		fmt.Printf("xxx products : %v \n", products)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, products)
	})
}