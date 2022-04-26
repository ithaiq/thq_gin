package internal

import "reflect"

//BeanFactory 依赖注入工厂
type BeanFactory struct {
	beans []interface{}
}

func NewBeanFactory() *BeanFactory {
	bf := &BeanFactory{beans: make([]interface{}, 0)}
	bf.beans = append(bf.beans, bf)
	return bf
}

func (b *BeanFactory) setBean(beans ...interface{}) {
	b.beans = append(b.beans, beans...)
}

func (b *BeanFactory) GetBean(bean interface{}) interface{} {
	return b.getBean(reflect.TypeOf(bean))
}

func (b *BeanFactory) getBean(r reflect.Type) interface{} {
	for _, p := range b.beans {
		if r == reflect.TypeOf(p) {
			return p
		}
	}
	return nil
}

func (t *BeanFactory) Inject(v IClass) {
	vRef := reflect.ValueOf(v).Elem()
	vT := reflect.TypeOf(v).Elem()
	for i := 0; i < vRef.NumField(); i++ {
		f := vRef.Field(i)
		if !f.IsNil() || f.Kind() != reflect.Ptr {
			continue
		}
		if IsAnnotation(f.Type()) {
			f.Set(reflect.New(f.Type().Elem()))
			f.Interface().(Annotation).SetTag(vT.Field(i).Tag)
			t.inject(f.Interface())
			continue
		}

		if p := t.getBean(f.Type()); p != nil {
			// vRef.Field(0).Type() --> 指针 *GormAdapter
			// vRef.Field(0).Type().Elem() -->指针指向的对象 GormAdapter
			f.Set(reflect.New(f.Type().Elem()))
			f.Elem().Set(reflect.ValueOf(p).Elem())
		}
	}
}

func (b *BeanFactory) inject(object interface{}) {
	rV := reflect.ValueOf(object)
	if rV.Kind() == reflect.Ptr {
		rV = rV.Elem()
	}
	for i := 0; i < rV.NumField(); i++ {
		f := rV.Field(i)
		if f.Kind() != reflect.Ptr || !f.IsNil() {
			continue
		}
		if p := b.getBean(f.Type()); p != nil && f.CanInterface() {
			f.Set(reflect.New(f.Type().Elem()))
			f.Elem().Set(reflect.ValueOf(p).Elem())
		}
	}
}
