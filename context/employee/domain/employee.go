package domain

import (
	"bettertomorrow/common/util"
	"database/sql"
	"time"
)

type Employee struct {
	Id        uint      `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	CreatedAt time.Time `json:"created_at"`
}

func (e *Employee) ToEmployeeRaw() EmployeeRaw {
	return EmployeeRaw{
		e.Id,
		util.IfNullString(e.FirstName == "") (sql.NullString{}, sql.NullString{e.FirstName, true}),
		util.IfNullString(e.LastName == "") (sql.NullString{}, sql.NullString{e.LastName, true}),
		e.CreatedAt,
	}
}

/*************/

type EmployeeRaw struct {
	Id        uint
	FirstName sql.NullString `db:"first_name"`
	LastName  sql.NullString `db:"last_name" `
	CreatedAt time.Time      `db:"created_at" `
}

func (er *EmployeeRaw) ToEmployee() Employee {
	return Employee{
		er.Id,
		util.IfString(er.FirstName.Valid) (er.FirstName.String, ""),
		util.IfString(er.LastName.Valid) (er.LastName.String, ""),
		er.CreatedAt,	
	}
}

func (EmployeeRaw) TableName() string {
	return "employees"
}
