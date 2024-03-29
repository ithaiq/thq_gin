package classes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ithaiq/tgin/example/model"
	"github.com/ithaiq/tgin/internal"
)

type IndexClass struct {
	*internal.GormAdapter
	Age *internal.Value `prefix:"user.score"`
}

func NewIndexClass() *IndexClass {
	return &IndexClass{}
}

func (i *IndexClass) GetIndex(ctx *gin.Context) string {
	var id int
	i.DB.Table("t1").Select("id").Limit(1).Scan(&id)
	fmt.Println(id)
	return "ok"
}

func (i *IndexClass) GetModel(ctx *gin.Context) internal.Model {
	user := &model.UserModel{}
	internal.Error(ctx.BindUri(user))
	return user
}

func (i *IndexClass) Test(c *gin.Context) string {
	return "测试" + i.Age.String()
}

func (i *IndexClass) Build(t *internal.TGin) {
	t.Handle("GET", "/index", i.GetIndex)
	t.Handle("GET", "/test", i.Test)
	t.Handle("GET", "/model/:id", i.GetModel)
}
