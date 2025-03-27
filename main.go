package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Carrega todos os templates da pasta "templates"
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	// Rota para a página inicial com suporte a idiomas
	router.GET("/", func(c *gin.Context) {
		lang := c.Query("lang")
		if lang == "" {
			lang = "en" // inglês padrão
		}

		var templateName string
		switch lang {
		case "pt":
			templateName = "index_pt.html"
		case "es":
			templateName = "index_es.html"
		default:
			templateName = "index_en.html"
		}

		c.HTML(http.StatusOK, templateName, gin.H{
			"title": "go-website-tutorial",
		})
	})

	router.Run(":8080")
}
