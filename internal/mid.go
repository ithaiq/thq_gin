package internal

import "github.com/gin-gonic/gin"

type Mid interface {
	OnRequest(ctx *gin.Context) error
}
