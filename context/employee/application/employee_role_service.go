/**
 * Arificialy created struct to have dependecies disturbutted among many structs (more than 1:1 (ex.: 1 repository <-> 1 service))
 */

package application

import (
	"bettertomorrow/context/employee/domain"
	"bettertomorrow/context/employee/persistance"
	"sync"
)

type EmployeeRolesService interface {
	FindEmployeeWithRoles() ([]domain.EmployeeWithRoles, error)
}


type EmployeeRolesServiceImpl struct {
	 employeeRepository persistance.EmployeeRepository
	 roleRepository persistance.RoleRepository
 }

 /* -- singleton for DI -- */

var employeeRolesServiceInstance *EmployeeRolesServiceImpl
var onceForEmployeeRolesService sync.Once

func ProvideEmployeeRolesServiceImpl() *EmployeeRolesServiceImpl {
	onceForEmployeeRolesService.Do(func() {
		employeeRolesServiceInstance = &EmployeeRolesServiceImpl{
			employeeRepository: persistance.ProvideEmployeeRepositoryImpl(),
			roleRepository: persistance.ProvideRoleRepositoryImpl(),
		}
	})

	return employeeRolesServiceInstance
} 

 /* ---- */

 func (cpsImpl *EmployeeRolesServiceImpl) FindEmployeeWithRoles() ([]domain.EmployeeWithRoles, error) {
	 return cpsImpl.employeeRepository.FindWithRoles()
 }

 