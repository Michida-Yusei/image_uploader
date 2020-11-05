package main // import "server"

import (
  "server/handler"

  "github.com/gin-gonic/gin" // フレームワーク
  "github.com/gin-contrib/cors"
)

func main() {
    r := gin.Default()

    r.Use(cors.New(cors.Config{
        AllowOrigins: []string{"http://localhost:8080"},
        AllowMethods: []string{"GET", "POST", "DELETE", "OPTIONS"},
        AllowHeaders: []string{"*"},
    }))

    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message" : "hello",
        })
    })

    r.POST("/images", handler.Upload)

    r.DELETE("/images/:uuid", handler.Delete)

    r.Run(":8888")

}

