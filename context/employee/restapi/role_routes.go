package restapi

import (
	"net/http"

	"bettertomorrow/context/employee/application"
	"bettertomorrow/context/employee/domain"

	"github.com/labstack/echo/v4"
)

type RoleRouter struct {
	roleService application.RoleService
}

func instantiateRoleRouter(roleService application.RoleService) *RoleRouter {
	return &RoleRouter{roleService}
}

func (router *RoleRouter) AddRoutes(apiRoutes *echo.Group) {
	apiRoutes.GET("/roles", func(c echo.Context) error {
		roles, err := router.roleService.FindAllRoles()
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, roles)
	})

	apiRoutes.POST("/roles", func(c echo.Context) error {
		var newRole domain.Role
		err := c.Bind(&newRole)
		if err != nil {
			return err
		}
		router.roleService.CreateRole(&newRole)	

		return c.JSON(http.StatusOK, "OK")
	})
}
