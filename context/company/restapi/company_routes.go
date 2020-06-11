package restapi

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	// "github.com/pkg/errors"

	//"time"
	"bettertomorrow/context/company/application"
	"bettertomorrow/context/company/domain"
)

type CompanyRouter struct {
	companyService                    *application.CompanyServiceImpl
	companyServiceWithProductsService *application.CompanyProductsServiceImpl
}

//TODO: pass interface
func instantiateCompanyRouter(companyService *application.CompanyServiceImpl, companyServiceWithProductsService *application.CompanyProductsServiceImpl) *CompanyRouter {
	return &CompanyRouter{companyService, companyServiceWithProductsService}
}

func (cr *CompanyRouter) AddRoutes(apiRoutes *echo.Group) {
	apiRoutes.GET("/companies", func(c echo.Context) error {
		companies, err := cr.companyService.FindAllCompanies()
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, companies)
	})

	apiRoutes.GET("/companies/products", func(c echo.Context) error {
		companiesWithProdcut, err := cr.companyServiceWithProductsService.FindCompanyWithProducts()
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, companiesWithProdcut)
	})

	apiRoutes.POST("/companies", func(c echo.Context) error {
		newCompany := &domain.Company{}

		err := c.Bind(newCompany)
		if err != nil {
			return err
		}
		//TODO: add logger remove fmt
		fmt.Printf("new company: %v \n", newCompany)

		err = cr.companyService.CreateCompany(newCompany)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, "OK")
	})

	apiRoutes.PUT("/companies", func(c echo.Context) error {
		updatedCompany := &domain.Company{}

		err := c.Bind(updatedCompany)
		if err != nil {
			return err
		}
		fmt.Printf("update company: %v \n", updatedCompany)

		err = cr.companyService.UpdateCompany(updatedCompany)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, "OK")
	})

	apiRoutes.DELETE("/companies/:id", func(c echo.Context) error {
		idStr := c.Param("id")
		id, err := strconv.ParseUint(idStr, 10, 32)
		if err != nil {
			return err
		}
		fmt.Printf("delete company with id: %v \n", id)

		err = cr.companyService.DeleteCompany(uint(id))
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, "OK")
	})
}
