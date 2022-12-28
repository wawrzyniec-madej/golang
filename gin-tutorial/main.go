package main

import "gin-tutorial/kernel"

func main() {

	kernel := kernel.NewKernel("dev")
	kernel.Start()
}
