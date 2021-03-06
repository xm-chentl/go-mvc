package container

import (
	"fmt"
	"sync"
)

var (
	mt        sync.Mutex
	keyOfInst = make(map[string]interface{})
)

func Get(key string) interface{} {
	fmt.Println(keyOfInst)
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

func Has(key string) bool {
	mt.Lock()
	defer mt.Unlock()

	_, ok := keyOfInst[key]

	return ok
}
