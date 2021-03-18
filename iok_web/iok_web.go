package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})

	r.GET("/auth/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "auth.tmpl", gin.H{
			"title": "Authentication",
		})
	})

	r.Run()
}
