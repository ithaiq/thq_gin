package internal

import "github.com/gin-gonic/gin"

//Mid 中间件
type Mid interface {
	OnRequest(ctx *gin.Context) error
}
