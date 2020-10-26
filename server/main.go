package main // import "server"

import (
  "server/handler"

  "github.com/gin-contrib/cors" // gin 用のCORS設定パッケージ
  "github.com/gin-gonic/gin"    // フレームワーク
)

func main() {
    r := gin.Default()

    config = cors.Config{
      AllowOrigins: []string{"http://localhost:8080"},
      AllowMethods: []string{"GET", "POST", "DELETE", "OPTIONS"},
      AllowHeaders: []string{"*"},
    }

    corsObj = cors.New(config)

    r.Use(corsObj)

    r.POST("/images", handler.Upload)

    r.DELETE("/images:uuid", handler.Delete)

    r.Run(":8888")
}
