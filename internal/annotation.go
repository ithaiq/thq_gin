package internal

import (
	"fmt"
	"reflect"
	"strings"
)

//Annotation 注解封装
type Annotation interface {
	SetTag(tag reflect.StructTag)
	String() string
}

var AnnotationList []Annotation

func IsAnnotation(t reflect.Type) bool {
	for _, item := range AnnotationList {
		if reflect.TypeOf(item) == t {
			return true
		}
	}
	return false
}

func init() {
	AnnotationList = append(AnnotationList, new(Value))
}

//Value tag结构封装
type Value struct {
	tag         reflect.StructTag
	BeanFactory *BeanFactory
}

func (v *Value) SetTag(tag reflect.StructTag) {
	v.tag = tag
}

func (v *Value) String() string {
	prefix := v.tag.Get("prefix")
	if prefix == "" {
		return ""
	}
	if config := v.BeanFactory.GetBean(new(SysConfig)); config != nil {
		rV := GetConfigValue(config.(*SysConfig).Config, strings.Split(prefix, "."), 0)
		if rV != nil {
			return fmt.Sprintf("%v", rV)
		}
		return ""
	}
	return ""
}
