package main // import "server"

import (
  "server/handler"

  "github.com/gin-contrib/cors" // gin 用のCORS設定パッケージ
  "github.com/gin-gonic/gin"    // フレームワーク
)

func main() {
    r := gin.Default()


    r.GET("/", func(c *gin.Context) {
      c.JSON(200, gin.H{
        "message" : "ping",
      })
    })

    r.POST("/images", handler.Upload)

    r.DELETE("/images:uuid", handler.Delete)

    r.Run(":8888")
}
