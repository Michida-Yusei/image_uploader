package main // import "server"

import (
    "fmt"
    "github.com/gin-gonic/gin" // フレームワーク
    "net/http"
    "os"

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
            "message" : "ping",
        })
    })

    r.POST("/images", Upload)

    r.DELETE("/images/:uuid", Delete)

    r.Run(":8888")

}


/**
 * Upload upload files.
 * gin.Context にはアップロードされたfile達を処理するための、メソッドがいくつか生えている。
 */
func Upload(c *gin.Context) {
    form, _ := c.MultipartForm()
    files := form.File["file"]

    // uuid を所得
    uuid := c.PostForm("uuid")

    for _, file := range files {
        // ファイル名にuuidを仕込む
        err := c.SaveUploadedFile(file, "images/"+uuid+".png")
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
        }
    }

    c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
}

/**
 * Delete remove file.
 */
func Delete(c *gin.Context) {
    uuid := c.Param("uuid")

    err := os.Remove(fmt.Sprintf("images/%s.png", uuid))
    if err != nil {
        fmt.Println(err.Error())
        c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("id: %s is deleted!", uuid)})
}
