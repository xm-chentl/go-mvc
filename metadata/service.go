package metadata

import (
	"reflect"
)

// Service 服务
type Service struct {
	Instance     interface{}       // 实例
	Name         string            // 不区分大小写
	NameOfAction map[string]Action // 对就在方法 key 全小写
	Type         reflect.Type      // 类型
}
