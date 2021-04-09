package container

import (
	"sync"
)

var (
	mt        sync.Mutex
	keyOfInst = make(map[string]interface{})
)

func Get(key string) interface{} {
	if inst, ok := keyOfInst[key]; ok {
		return inst
	}

	return nil
}

func Set(key string, inst interface{}) {
	mt.Lock()
	defer mt.Unlock()

	keyOfInst[key] = inst
}

func SetMany(insts map[string]interface{}) {
	mt.Lock()
	defer mt.Unlock()

	for key, inst := range insts {
		keyOfInst[key] = inst
	}
}

func Has(key string) bool {
	mt.Lock()
	defer mt.Unlock()

	_, ok := keyOfInst[key]

	return ok
}
