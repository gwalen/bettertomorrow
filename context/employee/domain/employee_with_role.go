package domain

type EmployeeWithRoles struct {
	Employee Employee
	Role     []Role
}

type EmployeeWithRole struct {
	Employee Employee `db:"employees"`
	Role     Role     `db:"roles"`
}

/*************/

type EmployeeWithRoleRaw struct {
	EmployeeRaw EmployeeRaw `db:"employees"`
	RoleRaw     RoleRaw     `db:"roles"`
}

func (erRaw *EmployeeWithRoleRaw) ToEmployeWithRole() EmployeeWithRole {
	return EmployeeWithRole{
		erRaw.EmployeeRaw.ToEmployee(),
		erRaw.RoleRaw.ToRole(),
	}
}

func (EmployeeWithRoleRaw) TableName() string {
	return "employees"
}
