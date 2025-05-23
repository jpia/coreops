package memcore

import (
	"coreops/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchCustomer(t *testing.T) {
	adapter := NewMemCoreAdapter()

	c1 := domain.Customer{SSN: "111-11-1111", FirstName: "Jane", LastName: "Smith"}
	c2 := domain.Customer{SSN: "222-22-2222", FirstName: "John", LastName: "Smythe"}
	_ = adapter.CreateCustomer(c1)
	_ = adapter.CreateCustomer(c2)

	results, err := adapter.SearchCustomer("smi")
	assert.NoError(t, err)
	assert.Len(t, results, 1)
}
