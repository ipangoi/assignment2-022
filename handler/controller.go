package handler

import "github.com/gin-gonic/gin"

type OrderHandler interface {
	GetAllOrder(*gin.Context)
	GetOneOrder(*gin.Context)
	CreateOrder(*gin.Context)
	UpdateOrder(*gin.Context)
	DeleteOrder(*gin.Context)
}
