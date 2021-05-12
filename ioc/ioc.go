package ioc

import (
	"reflect"
	"sync"
)

var (
	mt            sync.Mutex
	keyOfInstance = make(map[reflect.Type]interface{})
)

func Get(typeKey reflect.Type) interface{} {
	if inst, ok := keyOfInstance[typeKey]; ok {
		return inst
	}

	return nil
}

func Set(inst interface{}) {
	mt.Lock()
	defer mt.Unlock()

	keyOfInstance[reflect.TypeOf(inst)] = inst
}

func SetMany(instances ...interface{}) {
	mt.Lock()
	defer mt.Unlock()

	for _, instance := range instances {
		keyOfInstance[reflect.TypeOf(instance)] = instance
	}
}

func Has(typeKey reflect.Type) bool {
	mt.Lock()
	defer mt.Unlock()

	_, ok := keyOfInstance[typeKey]

	return ok
}
