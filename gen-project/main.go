package main

import "fmt"

func main() {

	PrintAnything("qwerty")

}

func PrintString(s string) {
	fmt.Print(s)
}

func PrintAnything[T int | string](s T) {
	fmt.Print(s)
}
