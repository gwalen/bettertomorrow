package domain

type CompanyWithProducts struct {
	Company Company
	Product []Product
}

type CompanyWithProduct struct {
	Company Company `xorm:"extends"`
	Product Product `xorm:"extends"`
}

func (CompanyWithProduct) TableName() string {
	return "companies"
}
