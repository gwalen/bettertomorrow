package domain

type CustomerWithWallets struct {
	Customer Customer
	Wallet   []Wallet
}

//`db` tag is for sqlx
type CustomerWithWallet struct {
	Customer Customer `xorm:"extends" db:"customers"`
	Wallet   Wallet   `xorm:"extends" db:"wallets"`
}

func (CustomerWithWallet) TableName() string {
	return "customers"
}