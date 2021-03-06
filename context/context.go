package context

import (
	"github.com/xm-chentl/go-mvc"
	"github.com/xm-chentl/go-mvc/enum"
)

// Context 上下文内容
type context map[enum.ContextValue]interface{}

// Get 获取
func (c context) Get(key enum.ContextValue) interface{} {
	return c[key]
}

// Has 是否存在
func (c context) Has(keys ...enum.ContextValue) bool {
	isOk := true
	for _, key := range keys {
		if _, ok := c[key]; !ok {
			isOk = ok
			break
		}
	}

	return isOk
}

// Set 设置
func (c context) Set(key enum.ContextValue, value interface{}) {
	c[key] = value
}

func New() mvc.IContext {
	return &context{}
}
