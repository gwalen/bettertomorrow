package application

import (
	"bettertomorrow/context/employee/domain"
	"bettertomorrow/context/employee/persistance"
	"sync"
)

type RoleService interface {
	CreateRole(employee *domain.Role)
	UpdateRole(employee *domain.Role)
	DeleteRole(id uint)
	FindAllRoles() ([]domain.Role, error)
}

type RoleServiceImpl struct {
	roleRepository persistance.RoleRepository
}

/* -- singleton for DI -- */

var roleServiceInstance *RoleServiceImpl
var onceForRoleService sync.Once

func ProvideRoleServiceImpl() *RoleServiceImpl {
	onceForRoleService.Do(func() {
		var roleRepository persistance.RoleRepository
		roleRepositoryImpl := persistance.ProvideRoleRepositoryImpl()
		roleRepository = roleRepositoryImpl
		roleServiceInstance = &RoleServiceImpl{roleRepository}
	})
	return roleServiceInstance
}

/* ---- */

func (impl *RoleServiceImpl) CreateRole(employee *domain.Role) {
	impl.roleRepository.Insert(employee)
}

func (impl *RoleServiceImpl) UpdateRole(employee *domain.Role) {
	impl.roleRepository.InsertOrUpdate(employee)
}

func (impl *RoleServiceImpl) DeleteRole(id uint) {
	impl.roleRepository.Delete(id)
}

func (impl *RoleServiceImpl) FindAllRoles() ([]domain.Role, error) {
	roles, err := impl.roleRepository.FindAll()
	return roles, err
}
