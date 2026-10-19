package main

import (
	"log"
	"net/http"
	"url-shortner/config"
	"url-shortner/service"

	"github.com/gin-gonic/gin"
)


type PostURLReq struct{
	LongURL string `json:"long_url" binding:"required,url"`
}
func main() {
  // Create a Gin router with default middleware (logger and recovery)
  r := gin.Default()

  rdb:=config.NewRedisClient()
  redisClient:=rdb.Client
  service:=service.NewService(redisClient)
  r.POST("/shorten",func (c *gin.Context){
	var req PostURLReq
	ctx:=c.Request.Context()
	err:=c.ShouldBind(&req)
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
	shortURL,err:=service.ShortenURL(ctx,req.LongURL)
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
	c.JSON(http.StatusOK ,gin.H{"short_url":shortURL})
  })

  r.GET("/:shortURL",func(c *gin.Context){
	shortURL:=c.Param("shortURL")
	longURL,err:=service.GetLongURL(c.Request.Context(),shortURL)
	if err!=nil{
		c.JSON(http.StatusNotFound,gin.H{"error":"URL not found"})
	}
	if longURL==""{
		c.JSON(http.StatusNotFound,gin.H{"error":"URL not found"})
		return
	}
	c.Redirect(http.StatusMovedPermanently,longURL)
  })
  // Define a simple GET endpoint
  r.GET("/ping", func(c *gin.Context) {
    // Return JSON response
    c.JSON(http.StatusOK, gin.H{
      "message": "pong",
    })
  })

 
  // Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
  if err := r.Run(); err != nil {
    log.Fatalf("failed to run server: %v", err)
  }
}