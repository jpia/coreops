package domain

type Customer struct {
    SSN       string `json:"ssn"`
    FirstName string `json:"first_name"`
    LastName  string `json:"last_name"`
}
