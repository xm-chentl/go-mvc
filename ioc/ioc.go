package ioc

import (
	"fmt"
	"reflect"
)

var (
	keyOfInstance = make(map[reflect.Type]reflect.Value)
)

func Get(inst interface{}) interface{} {
	instRt := getType(inst)
	if instRv, ok := keyOfInstance[instRt]; ok {
		return instRv.Interface()
	}

	return nil
}

func Set(instType interface{}, inst interface{}) {
	rt := getType(instType)
	instRt := reflect.TypeOf(inst)
	if !instRt.Implements(rt) {
		panic("Inst is not an InstType derived class")
	}
	keyOfInstance[rt] = reflect.ValueOf(inst)
}

func Has(inst interface{}) bool {
	instRt := getType(inst)
	_, ok := keyOfInstance[instRt]

	return ok
}

func getType(inst interface{}) reflect.Type {
	instRt, ok := inst.(reflect.Type)
	if !ok {
		instRt = reflect.TypeOf(inst)
	}
	if instRt.Kind() == reflect.Ptr {
		instRt = instRt.Elem()
	}
	if instRt.Kind() != reflect.Interface {
		panic(
			fmt.Errorf("inst is not interface"),
		)
	}

	return instRt
}
