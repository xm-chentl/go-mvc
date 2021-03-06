package metadata

import (
	"sync"

	"github.com/xm-chentl/go-mvc"
)

var (
	mt       sync.Mutex
	keyOfAPI = make(map[string]mvc.IApi)
)

// Register 注册
func Register(apis ...mvc.IApi) {
	mt.Lock()
	defer mt.Unlock()

	for _, api := range apis {
		keyOfAPI[api.Code()] = api
	}
}

// Get 获取Api
func Get(key string) mvc.IApi {
	mt.Lock()
	defer mt.Unlock()

	if inst, ok := keyOfAPI[key]; ok {
		return inst
	}

	return nil
}

// Has 判断是否存在
func Has(key string) bool {
	mt.Lock()
	defer mt.Unlock()

	_, ok := keyOfAPI[key]
	return ok
}
