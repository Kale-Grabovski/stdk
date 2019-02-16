package domain

import "github.com/gin-gonic/gin"

type Action interface {
	Handle(ctx *gin.Context)
}
