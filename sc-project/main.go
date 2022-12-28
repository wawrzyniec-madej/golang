package main

import (
	"sc-project/factories"
	"sc-project/services"
)

func main() {

	serviceContainer := services.NewServiceContainer()

	factory := factories.TestFactory{}

	serviceContainer.RegisterService("f", &factory)

	testFactory, ok := (*serviceContainer.GetServiceById("f")).(*factories.TestFactory)
	if !ok {
		panic("This is not valid struct!")
	}

	testFactory.Test()
}
