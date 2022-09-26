package core

type Seller struct {
	ID          int    `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Patronymic  string `json:"patronymic"`
	CompanyName string `json:"company_name"`
	ProductID   int    `json:"productID"`
}
