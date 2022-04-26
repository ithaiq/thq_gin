package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type UserMid struct{}

func NewUserMid() *UserMid {
	return &UserMid{}
}

func (this *UserMid) OnRequest(ctx *gin.Context) error {
	fmt.Println("这是新的用户中间件")
	return nil
}
