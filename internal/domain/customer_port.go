package domain

type CustomerPort interface {
	SearchCustomer(lastName string) ([]Customer, error)
	CreateCustomer(c Customer) error
}
