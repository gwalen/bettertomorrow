package domain

import (
	"time"
)

type Company struct {
	Address   `xorm:"extends" json:"address"`
	ID        uint      `gorm:"primary_key;AUTO_INCREMENT" xorm:"pk autoincr 'id'" json:"id"`
	Name      string    `json:"name"`
	TaxId     string    `json:"tax_id"`
	CreatedAt time.Time `xorm:"created" json:"-"`
	Products []Product  `xorm:"-"`
}

func (c Company) TableName() string {
	return "companies"
}


func (c *Company) Equals(c2 *Company) bool {
	return c.ID == c2.ID
}
