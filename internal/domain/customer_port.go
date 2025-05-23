package domain

type CustomerPort interface {
    FindCustomer(lastName string) ([]Customer, error)
    NewCustomer(c Customer) error
}
