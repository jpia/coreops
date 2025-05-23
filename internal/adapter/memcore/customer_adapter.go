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
    return &MemCoreAdapter{}
}

func (a *MemCoreAdapter) CreateCustomer(c domain.Customer) error {
    a.mu.Lock()
    defer a.mu.Unlock()
    a.customers = append(a.customers, c)
    return nil
}

func (a *MemCoreAdapter) SearchCustomer(lastName string) ([]domain.Customer, error) {
    a.mu.RLock()
    defer a.mu.RUnlock()
    var result []domain.Customer
    for _, c := range a.customers {
        if strings.Contains(strings.ToLower(c.LastName), strings.ToLower(lastName)) {
            result = append(result, c)
        }
    }
    return result, nil
}
