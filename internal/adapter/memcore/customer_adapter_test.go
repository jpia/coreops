package memcore_test

import (
    "testing"
    "coreops/internal/adapter/memcore"
    "coreops/internal/domain"
    "github.com/stretchr/testify/assert"
)

func TestMemCoreAdapter(t *testing.T) {
    adapter := memcore.NewMemCoreAdapter()

    c1 := domain.Customer{SSN: "111-11-1111", FirstName: "Jane", LastName: "Smith"}
    c2 := domain.Customer{SSN: "222-22-2222", FirstName: "John", LastName: "Smiley"}
    c3 := domain.Customer{SSN: "333-33-3333", FirstName: "Sara", LastName: "Smythe"}

    _ = adapter.NewCustomer(c1)
    _ = adapter.NewCustomer(c2)
    _ = adapter.NewCustomer(c3)

    results, err := adapter.FindCustomer("smi")
    assert.NoError(t, err)
    assert.Len(t, results, 2)
}
