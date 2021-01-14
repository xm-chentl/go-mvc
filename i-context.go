package mvc

import "github.com/xm-chentl/go-mvc/enum"

// IContext 上下文内容
type IContext interface {
	Get(key enum.Value) interface{}
	Has(keys ...enum.Value) bool
	Set(key enum.Value, value interface{})
}
