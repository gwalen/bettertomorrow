package domain

type Address struct {
	Street      string `json:"street"`
	HouseNumber string `json:"house_number"`
	PostalCode  string `json:"postal_code"`
	City        string `json:"city"`
	Country     string `json:"country"`
}
