package domain

import (
	"bettertomorrow/common/util"
	"database/sql"
	"time"
)

type Role struct {
	Id             uint      `json:"id"`
	Name           string    `json:"name"`
	DepartmentName string    `json:"department_name"`
	StartedAt      time.Time `json:"started_at"`
	FinishedAt     time.Time `json:"finished_at"`
	CreatedAt      time.Time `json:"-"`
	EmployeeId     uint      `json:"employee_id"`
}

func (r *Role) ToRoleRaw() RoleRaw {
	return RoleRaw{
		r.Id,
		r.Name,
		util.IfNullString(r.DepartmentName == "") (sql.NullString{}, sql.NullString{r.DepartmentName, true}),
		r.StartedAt,
		util.IfNullTime(r.FinishedAt == time.Time{}) (sql.NullTime{}, sql.NullTime{r.FinishedAt, true}),
		r.CreatedAt,
		r.EmployeeId,
	}
}

/*************/

type RoleRaw struct {
	Id             uint
	Name           string
	DepartmentName sql.NullString `db:"department_name"`
	StartedAt      time.Time      `db:"started_at"`
	FinishedAt     sql.NullTime   `db:"finished_at"`
	CreatedAt      time.Time      `db:"created_at"`
	EmployeeId     uint           `db:"employee_id"`
}

func (rw *RoleRaw) ToRole() Role {
	return Role{
		rw.Id,
		rw.Name,
		util.IfString(rw.DepartmentName.Valid) (rw.DepartmentName.String, ""),
		rw.StartedAt,
		util.IfTime(rw.FinishedAt.Valid) (rw.FinishedAt.Time, time.Time{}),
		rw.CreatedAt,
		rw.EmployeeId,
	}
}

func (RoleRaw) TableName() string {
	return "roles"
}