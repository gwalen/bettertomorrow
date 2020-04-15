package domain

import (
	"time"
)

type Product struct {
	ID        uint      `xorm:"pk 'id'" json:"id"`
	Name      string    `json:"name"`
	Price     float32   `json:"price"`
	CompanyID uint      `xorm:"'company_id'" json:"comapny_id"`
	CreatedAt time.Time `json:"-"`
}
