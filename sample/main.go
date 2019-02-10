package main

import(
  "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "Hello World!!!!!",
        })
    })

    r.GET("/get", func(c *gin.Context) {
      c.JSON(200, gin.H{
          "message": "Hello World Go!!!",
      })
    })

    r.POST("/post", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "post": "OK!!!",
        })
      })
    // ポートを設定しています。
    r.Run(":8080")
}