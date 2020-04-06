package restapi

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	//"time"
	"bettertomorrow/context/company/application"
	"bettertomorrow/context/company/domain"
)

type CompanyRouter struct {
	companyService *application.CompanyServiceImpl
}

func instantiateCompanyRouter(companyService *application.CompanyServiceImpl) *CompanyRouter {
	return &CompanyRouter{companyService}
}

/* -- singleton for DI -- */

//var companyRouterInstance *CompanyRouter
//var once sync.Once
//
//func ProvideCompanyRouterImpl() *CompanyRouter {
//	once.Do(func() {
//		companyService := application.ProvideCompanyServiceImpl()
//		companyRouterInstance = &CompanyRouter{companyService}
//	})
//	return companyRouterInstance
//}

/* ---- */

func (cr *CompanyRouter) AddRoutes(apiRoutes *echo.Group) {
	apiRoutes.GET("/companies", func(c echo.Context) error {
		//TODO: handle error
		companies, err := cr.companyService.FindAllCompanies()
		if err != nil {
			fmt.Errorf("error in retiving comapnies: %v", err)
		} 
		//TODO: json
		return c.JSON(http.StatusOK, companies)
	})

	apiRoutes.POST("/companies", func(c echo.Context) error {
		newCompany := &domain.Company{}

		err := c.Bind(newCompany)
		if err != nil {
			// fmt.Errorf("error reading company data from request: %v", err) //TODO: test it
			return err
		}
		fmt.Printf("new company: %v \n", newCompany)	

		err = cr.companyService.CreateCompany(newCompany)	
		//TODO: one method for this kind of error handling like in go-tools project
		if err != nil {
			// fmt.Errorf("error saving company: %v , error: %v", newCompany, err) //TODO: test it
			return err
		}

		return c.JSON(http.StatusOK, nil)
	})
}


