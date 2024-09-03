package service

import (
	"reflect"
)

var (
	services = make(map[interface{ ServiceInterface }]interface{})
)

func GetService[S interface{ ServiceInterface }](obj interface{ ServiceInterface }) S {
	service := services[obj]
	if service == nil {
		serviceType := reflect.TypeOf(obj)
		serviceRealType := serviceType.Elem()
		serviceValue := reflect.New(serviceRealType)
		serviceRealValue := serviceValue.Elem()
		serviceRealValueAddr := serviceRealValue.Addr()
		serviceInitMethod, _ := serviceType.MethodByName("OnInitService")
		serviceInitMethod.Func.Call([]reflect.Value{serviceRealValueAddr})

		service = serviceValue.Interface()
		services[obj] = service
	}
	return service.(S)
}

type Service struct{}

type ServiceInterface interface {
	OnInitService()
}
