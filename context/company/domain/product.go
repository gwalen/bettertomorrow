package domain

import (
	"time"
)

type Product struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Price     float32   `json:"price"`
	CompanyID uint      `json:"comapny_id"`
	CreatedAt time.Time `json:"-"`
}
