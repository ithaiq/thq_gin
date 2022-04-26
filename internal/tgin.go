package internal

import "github.com/gin-gonic/gin"

type TGin struct {
	*gin.Engine
	g *gin.RouterGroup
}

func NewTGin() *TGin {
	t := &TGin{Engine: gin.New()}
	t.Use(ErrorHandler())
	return t
}

func (t *TGin) Launch() {
	t.Run(":8080")
}

func (t *TGin) Mount(group string, cls ...IClass) *TGin {
	t.g = t.Group(group)
	for _, v := range cls {
		v.Build(t)
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
