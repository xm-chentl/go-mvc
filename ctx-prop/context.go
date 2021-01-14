package ctxprop

import "github.com/xm-chentl/go-mvc/enum"

// Context 上下文内容
type Context map[enum.Value]interface{}

// Get 获取
func (c Context) Get(key enum.Value) interface{} {
	return c[key]
}

// Has 是否存在
func (c Context) Has(keys ...enum.Value) bool {
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
func (c Context) Set(key enum.Value, value interface{}) {
	c[key] = value
}
