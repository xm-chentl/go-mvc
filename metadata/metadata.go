package metadata

import (
	"reflect"
	"strings"
	"sync"
)

var (
	rwm          sync.RWMutex
	keyOfService = make(map[string]Service)
)

// Register 注册元数据
func Register(instances ...interface{}) {
	rwm.Lock()
	defer rwm.Unlock()

	// 1. 反射获取相应类名，方法集合
	for _, instance := range instances {
		ser := toService(instance)
		keyOfService[ser.Name] = ser
	}
}

// RegisterByCustom 自定义注册
func RegisterByCustom(customOfMap map[string]interface{}) {
	rwm.RLock()
	defer rwm.RUnlock()

	for key, instance := range customOfMap {
		ser := toService(instance)
		keyOfService[strings.ToLower(key)] = ser
	}
}

// IsInitialize 是否初始化
func IsInitialize() bool {
	rwm.RLock()
	defer rwm.RUnlock()

	return len(keyOfService) > 0
}

// Get 获取服务
func Get(key string) Service {
	rwm.RLock()
	defer rwm.RUnlock()

	return keyOfService[key]
}

// Has 存在
func Has(keys ...string) bool {
	rwm.RLock()
	defer rwm.RUnlock()

	isOk := true
	for _, key := range keys {
		if _, ok := keyOfService[key]; !ok {
			isOk = ok
			break
		}
	}

	return isOk
}

func toService(instance interface{}) Service {
	rt := reflect.TypeOf(instance)
	name := strings.ToLower(rt.Name())

	return Service{
		Instance:     instance,
		Name:         name,
		NameOfAction: toActions(rt),
		Type:         rt,
	}
}

func toActions(serviceType reflect.Type) map[string]Action {
	res := make(map[string]Action)
	numOfMethod := serviceType.NumMethod()
	for i := 0; i < numOfMethod; i++ {
		method := serviceType.Method(i)
		name := strings.ToLower(method.Name)
		parameters := make([]Parameter, 0)
		for j := 0; j < method.Type.NumIn(); j++ {
			parameters = append(parameters, Parameter{
				Name: method.Type.Name(),
				Type: method.Type.In(j),
			})
		}
		res[name] = Action{
			Name:       name,
			Method:     method,
			Parameters: parameters,
		}
	}

	return res
}
