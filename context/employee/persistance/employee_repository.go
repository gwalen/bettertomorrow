package persistance

import (
	"bettertomorrow/common/dbsqlx"
	"bettertomorrow/context/employee/domain"
	"sync"
	"time"

	"github.com/jmoiron/sqlx"
)

type EmployeeRepository interface {
	Insert(employee *domain.Employee)
	Update(employee *domain.Employee)
	InsertOrUpdate(employee *domain.Employee)
	Delete(id uint)
	FindAll() ([]domain.Employee, error)
	FindWithRoles() ([]domain.EmployeeWithRoles, error)
}

type EmployeeRepositoryImpl struct {
	db *sqlx.DB
}

/* -- singleton for DI -- */

var employeeRepositoryInstance *EmployeeRepositoryImpl
var onceForEmployeeRepository sync.Once

func ProvideEmployeeRepositoryImpl() *EmployeeRepositoryImpl {
	onceForEmployeeRepository.Do(func() {
		dbHandle := dbsqlx.DB()
		employeeRepositoryInstance = &EmployeeRepositoryImpl{dbHandle}
	})
	return employeeRepositoryInstance
}

/* ---- */

func (impl *EmployeeRepositoryImpl) Insert(employee *domain.Employee) {
	employeeRaw := employee.ToEmployeeRaw()
	employeeRaw.Id = 0 // setting to 0 will trigger auto increment
	query := "INSERT INTO employees(id, first_name, last_name, created_at) VALUES (0, ?, ?, ?);"
	impl.db.MustExec(query, employeeRaw.FirstName, employeeRaw.LastName, time.Now().UTC()) //MustExec will panic on error (Exec will not)
	// no error returned it will panic on error
}

//this will only work on mysql (its how you do insertOnUpdate in Mysql with plain query)
func (impl *EmployeeRepositoryImpl) InsertOrUpdate(employee *domain.Employee){
	employeeRaw := employee.ToEmployeeRaw()
	query := "INSERT INTO employees(id, first_name, last_name, created_at) VALUES (0, ?, ?, ?) ON DUPLICATE KEY UPDATE first_name = ?, last_name = ?"
	impl.db.MustExec(query, employeeRaw.FirstName, employeeRaw.LastName, employeeRaw.FirstName, employeeRaw.LastName)
}

func (impl *EmployeeRepositoryImpl) Update(employee *domain.Employee) {
	employeeRaw := employee.ToEmployeeRaw()
	query := "UPDATE employees SET first_name = ?, last_name = ? where id = ?"
	impl.db.MustExec(query, employeeRaw.FirstName, employeeRaw.LastName)
}

func (impl *EmployeeRepositoryImpl) Delete(id uint) {
	query := "DELETE FROM employees WHERE id = ?"
	impl.db.MustExec(query, id)
}

//Select featches all rows to memory and populats the slice, for large data sets when paging is needed use Queryx or query with limits (or offsets)
func (impl *EmployeeRepositoryImpl) FindAll() ([]domain.Employee, error) {
	var employeesRaw []domain.EmployeeRaw
	err := impl.db.Select(&employeesRaw, "SELECT * FROM employees")

	employees :=  make([]domain.Employee, len(employeesRaw))
	for i, employeeRaw := range employeesRaw {
		employees[i] = employeeRaw.ToEmployee()
	}
	return employees, err
}

func (impl *EmployeeRepositoryImpl) FindWithRoles() ([]domain.EmployeeWithRoles, error) {
	var employeeWithRoleArr []domain.EmployeeWithRole

	query := `
		SELECT 
			employees.id "employees.id",
			employees.first_name "employees.first_name",
			employees.last_name "employees.last_name",
			employees.created_at "employees.created_at",
			roles.id "roles.id",
			roles.name "roles.name",
			roles.department_name "roles.department_name",
			roles.started_at "roles.started_at",
			roles.finished_at "roles.finished_at",
			roles.created_at "roles.created_at",
			roles.employee_id "roles.employee_id"
		FROM employees JOIN roles ON roles.employee_id = employees.id;
	`
	rows, err := impl.db.Queryx(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var erRaw domain.EmployeeWithRoleRaw
		err := rows.StructScan(&erRaw)
		if err != nil {
			return nil, err //most porbably some extral logging would be necessary here
		}
		employeeWithRoleArr = append(employeeWithRoleArr, erRaw.ToEmployeWithRole())
	}

	employeesWithRoles := mapJoin(employeeWithRoleArr)
	return employeesWithRoles, err
}

func mapJoin(employeeWithRoleArr []domain.EmployeeWithRole) []domain.EmployeeWithRoles {
	employeesWithRolesMap := make(map[domain.Employee][]domain.Role)
	var employeesWithRoles []domain.EmployeeWithRoles

	for _, elem := range employeeWithRoleArr {
		// no need to check if key exists coz appending element to nil slice creates a new slice with this elemnent	=> append(nil, elem.Role) == []domain.Role{elem.Role}
		employeesWithRolesMap[elem.Employee] = append(employeesWithRolesMap[elem.Employee], elem.Role)
	}

	for employee, rolees := range employeesWithRolesMap {
		newemployeeWithRoles := domain.EmployeeWithRoles{employee, rolees}
		employeesWithRoles = append(employeesWithRoles, newemployeeWithRoles)
	}

	return employeesWithRoles
}
