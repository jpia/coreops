package handler

import (
    "net/http"
    "coreops/internal/domain"
    "github.com/gin-gonic/gin"
)

type CustomerHandler struct {
    adapters map[string]domain.CustomerPort
}

func NewCustomerHandler(adapters map[string]domain.CustomerPort) *CustomerHandler {
    return &CustomerHandler{adapters: adapters}
}

func (h *CustomerHandler) CreateCustomer(c *gin.Context) {
    coreAdapter := c.Param("core_adapter")

    var customer domain.Customer
    if err := c.ShouldBindJSON(&customer); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    adapter, ok := h.adapters[coreAdapter]
    if !ok {
        c.JSON(http.StatusNotFound, gin.H{"error": "Unknown core adapter"})
        return
    }

    if err := adapter.NewCustomer(customer); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save customer"})
        return
    }

    c.Status(http.StatusCreated)
}

func (h *CustomerHandler) SearchCustomers(c *gin.Context) {
    coreAdapter := c.Param("core_adapter")
    lastName := c.Query("last_name")

    adapter, ok := h.adapters[coreAdapter]
    if !ok {
        c.JSON(http.StatusNotFound, gin.H{"error": "Unknown core adapter"})
        return
    }

    results, err := adapter.FindCustomer(lastName)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Search failed"})
        return
    }

    c.JSON(http.StatusOK, results)
}
