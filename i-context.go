package mvc

import "github.com/xm-chentl/go-mvc/enum"

// IContext mvc上下文内容
type IContext interface {
	Set(key enum.ContextValue, value interface{})
	Get(key enum.ContextValue) interface{}
	Has(keys ...enum.ContextValue) bool
}
