package domain

import (
	"time"
)

//`db` tag is for sqlx
type Customer struct {
	Id        uint      `xorm:"pk autoincr" json:"id"`
	FirstName string    `db:"first_name" json:"first_name"`
	LastName  string    `db:"last_name" json:"last_name"`
	CreatedAt time.Time `db:"created_at" xorm:"created" json:"-"`
}

func (c Customer) TableName() string {
	return "customers"
}

func (c *Customer) Equals(c2 *Customer) bool {
	return c.Id == c2.Id
}
