package main

import (
	"etapa2/i18n"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.SetFuncMap(template.FuncMap{
		"safe": func(s string) template.HTML { return template.HTML(s) },
	})
	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		lang := c.DefaultQuery("lang", "en")

		texts, ok := i18n.Translations[lang]
		if !ok {
			texts = i18n.Translations["en"]
		}

		c.HTML(http.StatusOK, "index.html", texts)
	})

	router.Run(":8080")
}
