package handler

import (
  "fmt"
  "net/http"
  "os"

  "github.com/gin-gonic/gin"
)

/**
 * Upload upload files.
 * gin.Context にはアップロードされたfile達を処理するための、メソッドがいくつか生えている。
 */
func Upload(c *gin.Context) {
  form, _ := c.MultipartForm()
  files := form.File["file"]

  for _, file := range files {
    err := c.SaveUploadedFile(file, "images/" + file.Filename)
    if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
    }
  }

  c.JSON(http.StatusOK, gin.H{"message": "success!!"})
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
