package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	v1 := router.Group("/api/sessions")
	{
		v1.POST("/", CreateSession)
		v1.GET("/", GetSessions)
		v1.GET("/:session_id", GetSession)
		v1.DELETE("/:session_id", DeleteSession)
	}

	router.Run()

}