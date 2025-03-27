package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    router.GET("/", func(c *gin.Context) {
        c.String(http.StatusOK, "Hello, World! Bem-vindo ao projeto go-website-tutorial!")
    })

    router.Run(":8080")
}
