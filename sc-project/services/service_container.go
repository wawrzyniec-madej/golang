package services

import (
	"fmt"
)

type serviceContainer struct {
	services map[string]any
}

func NewServiceContainer() *serviceContainer {
	return &serviceContainer{
		services: map[string]any{},
	}
}

func (sc *serviceContainer) GetServices() map[string]any {
	return sc.services
}

func (sc *serviceContainer) RegisterService(id string, pointer any) {

	if sc.services[id] != nil {
		panic(fmt.Sprintf("Service '%s' is already registered!", id))
	}

	sc.services[id] = pointer
}

func (sc *serviceContainer) GetServiceById(id string) *any {
	pointerToService := sc.services[id]

	if pointerToService == nil {
		panic(fmt.Sprintf("Service '%s' does not exist!", id))
	}

	return &pointerToService
}
