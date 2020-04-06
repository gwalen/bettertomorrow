package domain

import "time"

type Company struct {
	Address   `json:"address"`
	ID        uint32 `gorm:"primary_key;AUTO_INCREMENT" json:"name"`
	Name      string `json:"id"`
	TaxId     string `json:"tax_id"`
	CreatedAt time.Time
}
