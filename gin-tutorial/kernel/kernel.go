package kernel

import (
	"gin-tutorial/routes"

	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slices"
)

type kernel struct {
	env string
}

func NewKernel(env string) *kernel {

	envs := []string{"prod", "dev"}

	if !slices.Contains(envs, env) {
		panic("This env is not supported")
	}

	return &kernel{
		env: env,
	}
}

func (kernel *kernel) Start() {
	router := gin.Default()
	routes.RegisterRoutes(router)
	router.Run("localhost:8080")
}
