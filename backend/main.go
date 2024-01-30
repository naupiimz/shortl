package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/naupiimz/shortl/backend/handler"
	"github.com/naupiimz/shortl/backend/store"
  "github.com/gin-contrib/cors"
)
func main() {
  r:= gin.Default()
  r.Use(cors.New(cors.Config{
    AllowOrigins: []string{"*"},
    AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
    AllowHeaders: []string{"Content-Type", "Authorization"},
  }))
  r.GET("/", func(ctx *gin.Context) {
    ctx.JSON(200, gin.H{
      "message":"Hey shortl",
    })
  })
  
  r.POST("/create-short-url", func(ctx *gin.Context) {
    handler.CreateShortUrl(ctx)
  })

  r.GET("/:shortUrl", func(ctx *gin.Context) {
    handler.HandleShortUrlRedirect(ctx)
  })

  store.InitializeStore()

  err:= r.Run(":9000")
  if err != nil{
    panic(fmt.Sprintf("Failed to start the web server on port 9000 - Error %v", err))
  }
}
