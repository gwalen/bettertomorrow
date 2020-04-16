package domain

import (
	"time"
)

type Company struct {
	Address   `xorm:"extends" json:"address"`
	ID        uint      `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Name      string    `json:"name"`
	TaxId     string    `json:"tax_id"`
	CreatedAt time.Time `json:"-"`
	Products []Product
}

func (c *Company) Equals(c2 *Company) bool {
	return c.ID == c2.ID
}
