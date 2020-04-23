package domain

import (
	"time"
)

type Customer struct {
	Id        uint      `xorm:"pk autoincr" json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	CreatedAt time.Time `xorm:"created" json:"-"`
}

func (c Customer) TableName() string {
	return "customers"
}

func (c *Customer) Equals(c2 *Customer) bool {
	return c.Id == c2.Id
}
