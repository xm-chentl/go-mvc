package mvc

import (
	"sync"
)

// DEFAULT 默认
const DEFAULT = "default"

var (
	rwm       sync.RWMutex
	nameOfMvc = make(map[string]IMvc)
)

// Default 默认
func Default() IMvc {
	rwm.RLock()
	defer rwm.RUnlock()

	ins, ok := nameOfMvc[DEFAULT]
	if !ok {
		panic("mvc is not instance")
	}

	return ins
}

// SetDefault 设置默认
func SetDefault(ins IMvc) {
	rwm.Lock()
	defer rwm.Unlock()

	nameOfMvc[DEFAULT] = ins
}
