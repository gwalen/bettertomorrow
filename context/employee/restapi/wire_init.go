// +build wireinject

package restapi

import (
	"bettertomorrow/context/employee/application"
	"github.com/google/wire"
)

func NewEmployeeRouter() *EmployeeRouter {
	wire.Build(
		application.ProvideEmployeeServiceImpl,
		application.ProvideEmployeeRolesServiceImpl,
		wire.Bind(new(application.EmployeeService), new(*application.EmployeeServiceImpl)),
		wire.Bind(new(application.EmployeeRolesService), new(*application.EmployeeRolesServiceImpl)),
		instantiateEmployeeRouter,
	)
	return &EmployeeRouter{}
}

func NewRoleRouter() *RoleRouter {
	wire.Build(
		application.ProvideRoleServiceImpl,
		wire.Bind(new(application.RoleService), new(*application.RoleServiceImpl)),
		instantiateRoleRouter,
	)
	return &RoleRouter{}
}
