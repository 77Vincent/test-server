package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"test-server/internal"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("internal/templates/*")
	r.Static("/assets", "./internal/assets")

	r.GET("/", func(c *gin.Context) {

		c.HTML(http.StatusOK, "index.gohtml", gin.H{
			"title": "Little gallery",
			"blogs": internal.Blogs,
		})
	})

	r.Run()
}
