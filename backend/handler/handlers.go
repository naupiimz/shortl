package handler

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/naupiimz/shortl/backend/shortener"
	"github.com/naupiimz/shortl/backend/store"
)

type UrlCreationRequest struct{
  LongUrl string `json:"long_url" binding:"required"`
  UserId string `json:"user_id" binding:"required"`
}

func CreateShortUrl(c *gin.Context)  {
  err:= godotenv.Load()
  if err != nil{
    log.Fatal("Error loading .env file")
  }

  host_url := os.Getenv("HOST_URL")
  
  var creationRequest UrlCreationRequest
  if err := c.ShouldBindJSON(&creationRequest); err != nil{
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
  return
  }

  shortUrl := shortener.GenerateShortLink(creationRequest.LongUrl, creationRequest.UserId)
  store.SaveUrlMapping(shortUrl, creationRequest.LongUrl, creationRequest.UserId)

  host:= host_url
  c.JSON(200, gin.H{
    "message":"short url created successfully",
    "short_url":host+shortUrl,
  })
}

func HandleShortUrlRedirect(c *gin.Context)  {
  shortUrl := c.Param("shortUrl")
  initialUrl:= store.RetrieveInitialUrl(shortUrl)
  c.Redirect(302, initialUrl)
}
