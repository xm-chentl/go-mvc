package ioc

import (
	"fmt"
	"reflect"
)

const (
	// Inject 注入
	TagInject = "inject"
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

func Has(inst interface{}) bool {
	instRt := getType(inst)
	_, ok := keyOfInstance[instRt]

	return ok
}

func Inject(structInst interface{}, funcs ...func(reflect.StructField) interface{}) error {
	rt := reflect.TypeOf(structInst).Elem()
	rv := reflect.ValueOf(structInst).Elem()
	for i := 0; i < rv.NumField(); i++ {
		field := rt.Field(i)
		if field.Type.Kind() == reflect.Interface {
			_, ok := field.Tag.Lookup(TagInject)
			if ok {
				if Has(field.Type) {
					// desc: 注入组件
					inst := Get(field.Type)
					if inst == nil {
						return fmt.Errorf("ioc inject faild err: %s is nil", field.Name)
					}
					rv.Field(i).Set(
						reflect.ValueOf(inst),
					)
					continue
				}
			}

			if len(funcs) > 0 {
				inst := funcs[0](field)
				if inst != nil {
					rv.Field(i).Set(
						reflect.ValueOf(inst),
					)
				}
			}
		}
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
