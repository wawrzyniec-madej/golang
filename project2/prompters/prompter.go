package prompters

import "fmt"

type Prompter struct{}

func (p Prompter) Prompt(message string) string {
	value := ""

	fmt.Printf("%s: ", message)
	fmt.Scan(&value)

	return value
}
