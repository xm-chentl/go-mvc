package mvc

import "sync"

const DEFAULT = "default"

var (
	mt       sync.Mutex
	keyOfMvc = make(map[string]IMvc)
)

func Default() IMvc {
	mt.Lock()
	defer mt.Unlock()

	if inst, ok := keyOfMvc[DEFAULT]; ok {
		return inst
	}
	panic("default instance is not exist")
}

func SetDefault(inst IMvc) {
	if inst == nil {
		return
	}

	mt.Lock()
	defer mt.Unlock()

	keyOfMvc[DEFAULT] = inst
}
