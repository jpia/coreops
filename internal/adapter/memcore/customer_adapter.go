package memcore

import (
    "strings"
    "sync"
    "coreops/internal/domain"
)

type MemCoreAdapter struct {
    mu        sync.RWMutex
    customers []domain.Customer
}

func NewMemCoreAdapter() *MemCoreAdapter {
    return &MemCoreAdapter{
        customers: make([]domain.Customer, 0),
    }
}

func (a *MemCoreAdapter) NewCustomer(c domain.Customer) error {
    a.mu.Lock()
    defer a.mu.Unlock()

    a.customers = append(a.customers, c)
    return nil
}

func (a *MemCoreAdapter) FindCustomer(lastName string) ([]domain.Customer, error) {
    a.mu.RLock()
    defer a.mu.RUnlock()

    var results []domain.Customer
    for _, c := range a.customers {
        if strings.Contains(strings.ToLower(c.LastName), strings.ToLower(lastName)) {
            results = append(results, c)
        }
    }

    return results, nil
}
