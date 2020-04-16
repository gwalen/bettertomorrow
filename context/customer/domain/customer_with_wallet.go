package domain

type CustomerWithWallets struct {
	Customer Customer
	Wallet   []Wallet
}

type CustomerWithWallet struct {
	Customer Customer `xorm:"extends"`
	Wallet   Wallet   `xorm:"extends"`
}

func (CustomerWithWallet) TableName() string {
	return "customers"
}
