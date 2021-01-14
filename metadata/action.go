package metadata

import (
	"reflect"
)

// Action 功能
type Action struct {
	Name       string         // 方法名
	Method     reflect.Method // 方法(函数)
	Parameters []Parameter
}
