package metadata

import (
	"github.com/xm-chentl/go-mvc"
)

var (
	keyOfAPI = make(map[string]mvc.IApi)
)

// Register 注册
func Register(key string, api mvc.IApi) {
	keyOfAPI[key] = api
}

// RegisterMap 使用键值注册
func RegisterMap(data map[string]mvc.IApi) {
	for k, v := range data {
		keyOfAPI[k] = v
	}
}

// Get 获取Api
func Get(key string) mvc.IApi {
	if inst, ok := keyOfAPI[key]; ok {
		return inst
	}

	return nil
}

// Has 判断是否存在
func Has(key string) bool {
	_, ok := keyOfAPI[key]
	return ok
}
