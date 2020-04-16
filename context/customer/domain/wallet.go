package domain

import (
	"time"
)

// in xorm evolves fied db names by converting camel case to snake case also for field with name ID which in Gorm is a specialy treated one (as private key)
type Wallet struct {
	Id         uint      `xorm:"pk" json:"id"`
	Amount     float32   `json:"amount"`
	Currency   string    `json:"currency"`
	CustomerId uint      `json:"customer_id"`
	CreatedAt  time.Time `json:"-"`
}
