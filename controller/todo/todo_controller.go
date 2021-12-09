package controller

import "github.com/gin-gonic/gin"

type TodoController interface {
	GetAll(c *gin.Context)
	GetById(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}
