package internal

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"reflect"
)

type TGin struct {
	*gin.Engine
	g     *gin.RouterGroup
	props []interface{}
}

func NewTGin() *TGin {
	t := &TGin{Engine: gin.New()}
	t.Use(ErrorHandler())
	return t
}

func (t *TGin) Launch() {
	config := InitConfig()
	t.Run(fmt.Sprintf(":%d", config.Server.Port))
}

func (t *TGin) Mount(group string, cls ...IClass) *TGin {
	t.g = t.Group(group)
	for _, v := range cls {
		v.Build(t)
		t.setProp(v)
	}
	return t
}

func (t *TGin) Handle(httpMethod, relativePath string, handler interface{}) *TGin {
	if h := Convert(handler); h != nil {
		t.g.Handle(httpMethod, relativePath, h)
	}
	return t
}

func (t *TGin) Attach(f Mid) *TGin {
	t.Use(func(c *gin.Context) {
		err := f.OnRequest(c)
		if err != nil {
			c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		} else {
			c.Next()
		}
	})
	return t
}

func (t *TGin) Beans(beans ...interface{}) *TGin {
	t.props = append(t.props, beans...)
	return t
}

func (t *TGin) getProp(r reflect.Type) interface{} {
	for _, p := range t.props {
		if r == reflect.TypeOf(p) {
			return p
		}
	}
	return nil
}

func (t *TGin) setProp(v IClass) {
	vRef := reflect.ValueOf(v).Elem()
	for i := 0; i < vRef.NumField(); i++ {
		f := vRef.Field(i)
		if !f.IsNil() || f.Kind() != reflect.Ptr {
			continue
		}
		if p := t.getProp(f.Type()); p != nil {
			// vRef.Field(0).Type() --> 指针 *GormAdapter
			// vRef.Field(0).Type().Elem() -->指针指向的对象 GormAdapter
			f.Set(reflect.New(f.Type().Elem()))
			f.Elem().Set(reflect.ValueOf(p).Elem())
		}
	}
}
