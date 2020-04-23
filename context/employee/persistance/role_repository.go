package persistance

import (
	"bettertomorrow/common/dbsqlx"
	"bettertomorrow/context/employee/domain"
	"sync"
	"time"

	"github.com/jmoiron/sqlx"
)

type RoleRepository interface {
	Insert(role *domain.Role)
	Update(role *domain.Role)
	InsertOrUpdate(role *domain.Role)
	Delete(id uint)
	FindAll() ([]domain.Role, error)
}

type RoleRepositoryImpl struct {
	db *sqlx.DB
}

/* -- singleton for DI -- */

var roleRepositoryInstance *RoleRepositoryImpl
var onceForRoleRepository sync.Once

func ProvideRoleRepositoryImpl() *RoleRepositoryImpl {
	onceForRoleRepository.Do(func() {
		dbHandle := dbsqlx.DB()
		roleRepositoryInstance = &RoleRepositoryImpl{dbHandle}
	})
	return roleRepositoryInstance
}

/* ---- */

func (impl *RoleRepositoryImpl) Insert(role *domain.Role) {
	roleRaw := role.ToRoleRaw()
	query := "INSERT INTO roles(id, name, department_name, started_at, finished_at, created_at, employee_id) VALUES (0, ?, ?, ?, ?, ?, ?)"
	impl.db.MustExec(query, roleRaw.Name, roleRaw.DepartmentName, roleRaw.StartedAt, roleRaw.FinishedAt, time.Now().UTC(), roleRaw.EmployeeId)
	// no error returned it will panic on error
}

func (impl *RoleRepositoryImpl) InsertOrUpdate(role *domain.Role) {
	query := `
		INSERT INTO roles(id, name, department_name, started_at, finished_at, created_at, employee_id) 
		VALUES (0, ?, ?, ?, ?, ?, ?)
		ON DUPLICATE KEY UPDATE name = ?, department_name = ?, started_at = ?, finished_at = ?, created_at = ?, employee_id = ?"
	`
	impl.db.MustExec(
		query,
		role.Id, role.Name, role.DepartmentName, role.StartedAt, role.FinishedAt, role.CreatedAt, role.EmployeeId, role.Name, 		
		role.DepartmentName, role.StartedAt, role.FinishedAt, role.CreatedAt, role.EmployeeId,
	)
}

func (impl *RoleRepositoryImpl) Update(role *domain.Role) {
	query := "UPDATE roles SET name = ?, department_name = ?, started_at = ?, finished_at = ?, created_at = ?, employee_id = ? where id = ?"
	impl.db.MustExec(query, role.DepartmentName, role.StartedAt, role.FinishedAt, role.CreatedAt, role.EmployeeId, role.Id)
}

func (impl *RoleRepositoryImpl) Delete(id uint) {
	query := "DELETE FROM roles WHERE id = ?"
	impl.db.MustExec(query, id)
}

func (impl *RoleRepositoryImpl) FindAll() ([]domain.Role, error) {
	var rolesRaw []domain.RoleRaw
	err := impl.db.Select(&rolesRaw, "SELECT * FROM roles")

	roles := make([]domain.Role, len(rolesRaw))
	for i, roleRaw := range rolesRaw {
		// roles = append(roles, *roleRaw.ToRole())		
		roles[i] = roleRaw.ToRole()
	}

	return roles, err
}
