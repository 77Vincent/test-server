package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Image struct {
	Src         string
	Description string
}

type blog struct {
	Title  string
	Images []*Image
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("internal/templates/*")
	r.Static("/assets", "./internal/assets")

	r.GET("/", func(c *gin.Context) {

		c.HTML(http.StatusOK, "index.gohtml", gin.H{
			"title": "Little gallery",
			"blogs": []*blog{
				{
					Title: "We had a date at the PAM",
					Images: []*Image{
						{
							Src:         "https://vincent-1316174341.cos.ap-nanjing.myqcloud.com/website/us.jpeg",
							Description: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
						},
					},
				},
			},
		})
	})

	r.Run()
}
