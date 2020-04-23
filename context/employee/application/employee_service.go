package application

import (
	"bettertomorrow/context/employee/domain"
	"bettertomorrow/context/employee/persistance"
	"sync"
)

type EmployeeService interface {
	CreateEmployee(employee *domain.Employee)
	UpdateEmployee(employee *domain.Employee)
	DeleteEmployee(id uint)
	FindAllEmployees() ([]domain.Employee, error)
}

type EmployeeServiceImpl struct {
	employeeRepository persistance.EmployeeRepository
}

/* -- singleton for DI -- */

var employeeServiceInstance *EmployeeServiceImpl
var onceForCompantService sync.Once

func ProvideEmployeeServiceImpl() *EmployeeServiceImpl {
	onceForCompantService.Do(func() {
		employeeRepository := persistance.ProvideEmployeeRepositoryImpl()
		employeeServiceInstance = &EmployeeServiceImpl{employeeRepository}
	})
	return employeeServiceInstance
}

/* ---- */

func (impl *EmployeeServiceImpl) CreateEmployee(employee *domain.Employee) {
	impl.employeeRepository.Insert(employee)
}

func (impl *EmployeeServiceImpl) UpdateEmployee(employee *domain.Employee) {
	impl.employeeRepository.Update(employee)
}

func (impl *EmployeeServiceImpl) DeleteEmployee(id uint) {
	impl.employeeRepository.Delete(id)
}

func (impl *EmployeeServiceImpl) FindAllEmployees() ([]domain.Employee, error) {
	employees, err := impl.employeeRepository.FindAll()
	return employees, err
}
