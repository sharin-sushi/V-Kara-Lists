package main

import (
	"fmt"

	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sharin-sushi/0016go_next_relation/infra"

	// "github.com/sharin-sushi/0016go_next_relation/interfaces/controllers"
	"github.com/sharin-sushi/0016go_next_relation/interfaces/controllers/common"
)

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost", "https://localhost",
			"http://v-karaoke.com", "https://v-karaoke.com",
			"http://backend.v-karaoke.com", "https://backend.v-karaoke.com",
		},
		AllowMethods: []string{"POST", "GET", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{
			"Origin", "Content-Length", "Content-Type", "Cookie",
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Accept-Encoding",
			"Authorization",
			"access-control-allow-origin",
			"Access-Control-Allow-Origin",
		},
		AllowCredentials: true,
	}))

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{})
	})

	r.Use(requestLogger()) //開発用。本番稼動時はコメントアウトする。

	infra.Routing(r)

	env := common.ReturnEvnCloudorLocal()
	var host string
	if env == "on cloud" {
		//クラウド環境
		// fmt.Println("hallow api server in AWS" + controllers.UpdateData)
		// host = "v-karaoke.com"
	} else if env == "on local" {
		// ローカルのdocker上(compose使用) or  VSCodeで起動
		// ローカルのファイアーウォール対策
		host = "localhost"
	}
	r.Run(host + ":8080")
}

func requestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Printf("Method: %s, Path: %s, Header: %v\n, , Body: %v\n", c.Request.Method, c.Request.URL.Path, c.Request.Header, c.Request.Body)
		c.Next()
	}
}
