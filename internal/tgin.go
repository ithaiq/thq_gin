package internal

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type IClass interface {
	Build(t *TGin)
}

//TGin gin封装
type TGin struct {
	*gin.Engine
	g           *gin.RouterGroup
	beanFactory *BeanFactory
}

func NewTGin() *TGin {
	t := &TGin{Engine: gin.New(), beanFactory: NewBeanFactory()}
	t.Use(ErrorHandler())
	t.beanFactory.setBean(InitConfig())
	return t
}

func (t *TGin) Launch() {
	var port int32 = 8080
	if config := t.beanFactory.GetBean(new(SysConfig)); config != nil {
		port = config.(*SysConfig).Server.Port
	}
	getCronTask().Start()
	t.Run(fmt.Sprintf(":%d", port))
}

func (t *TGin) Mount(group string, cls ...IClass) *TGin {
	t.g = t.Group(group)
	for _, v := range cls {
		v.Build(t)
		t.beanFactory.Inject(v)
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
	t.beanFactory.setBean(beans...)
	return t
}

func (t *TGin) Task(expr string, f func()) *TGin {
	err := getCronTask().AddFunc(expr, f)
	if err != nil {
		panic(err)
	}
	return t
}
