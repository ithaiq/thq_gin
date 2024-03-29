package internal

import (
	"github.com/gin-gonic/gin"
	"reflect"
)

var ResponderList []Responder

func init() {
	ResponderList = []Responder{
		new(StringResponder),
		new(ModelResponder),
	}
}

//Responder 响应封装
type Responder interface {
	RespondTo() gin.HandlerFunc
}

type StringResponder func(ctx *gin.Context) string

func (this StringResponder) RespondTo() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.String(200, this(context))
	}
}

type ModelResponder func(ctx *gin.Context) Model

func (this ModelResponder) RespondTo() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(200, this(context))
	}
}

type ModelsResponder func(ctx *gin.Context) Models

func (this ModelsResponder) RespondTo() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Writer.Header().Set("Content-type", "application/json")
		context.Writer.WriteString(string(this(context)))
	}
}

func Convert(handler interface{}) gin.HandlerFunc {
	hRef := reflect.ValueOf(handler)
	for _, r := range ResponderList {
		rRef := reflect.ValueOf(r).Elem()
		if hRef.Type().ConvertibleTo(rRef.Type()) {
			rRef.Set(hRef)
			return rRef.Interface().(Responder).RespondTo()
		}
	}
	return nil
}
