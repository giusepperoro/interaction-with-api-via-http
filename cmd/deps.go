package main

import "github.com/gin-gonic/gin"

type addToDatabaseHandler interface {
	HandleAddToDatabase(c *gin.Context)
}
