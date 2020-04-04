package restapi

import (
	"net/http"

	"github.com/labstack/echo/v4"

	//"time"
	"bettertomorrow/context/company/application"
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
		companies, _ := cr.companyService.FindAllCompanies()
		//TODO: json
		return c.JSON(http.StatusOK, companies)
	})
}
