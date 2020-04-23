package restapi

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	// "github.com/pkg/errors"

	//"time"
	"bettertomorrow/context/employee/application"
	"bettertomorrow/context/employee/domain"
)

type EmployeeRouter struct {
	employeeService                 application.EmployeeService
	employeeServiceWithRolesService application.EmployeeRolesService
}

//TODO: pass interface
func instantiateEmployeeRouter(
	employeeService application.EmployeeService,
	employeeServiceWithRolesService application.EmployeeRolesService,
) *EmployeeRouter {
	return &EmployeeRouter{employeeService, employeeServiceWithRolesService}
}

func (cr *EmployeeRouter) AddRoutes(apiRoutes *echo.Group) {
	apiRoutes.GET("/employees", func(c echo.Context) error {
		employees, err := cr.employeeService.FindAllEmployees()
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, employees)
	})

	apiRoutes.GET("/employees/roles", func(c echo.Context) error {
		employeesWithProdcut, err := cr.employeeServiceWithRolesService.FindEmployeeWithRoles()
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, employeesWithProdcut)
	})

	apiRoutes.POST("/employees", func(c echo.Context) error {
		newEmployee := &domain.Employee{}
		err := c.Bind(newEmployee)
		if err != nil {
			return err
		}
		cr.employeeService.CreateEmployee(newEmployee)

		return c.JSON(http.StatusOK, "OK")
	})

	apiRoutes.PUT("/employees", func(c echo.Context) error {
		updatedEmployee := &domain.Employee{}
		err := c.Bind(updatedEmployee)
		if err != nil {
			return err
		}
		fmt.Printf("update employee: %v \n", updatedEmployee)

		cr.employeeService.UpdateEmployee(updatedEmployee)

		return c.JSON(http.StatusOK, "OK")
	})

	apiRoutes.DELETE("/employees/:id", func(c echo.Context) error {
		idStr := c.Param("id")
		id, err := strconv.ParseUint(idStr, 10, 32)
		if err != nil {
			return err
		}
		fmt.Printf("delete employee with id: %v \n", id)

		cr.employeeService.DeleteEmployee(uint(id))

		return c.JSON(http.StatusOK, "OK")
	})
}
