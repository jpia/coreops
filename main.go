package main

import (
    "coreops/internal/adapter/memcore"
    "coreops/internal/handler"
    "coreops/internal/domain"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    mem := memcore.NewMemCoreAdapter()

    adapterMap := map[string]domain.CustomerPort{
        "memcore": mem,
    }

    customerHandler := handler.NewCustomerHandler(adapterMap)

    r.POST("/:core_adapter/customer", customerHandler.CreateCustomer)
    r.GET("/:core_adapter/customer", customerHandler.SearchCustomers)

    r.Run(":8080")
}
