package factories

import "fmt"

type TestFactory struct{}

func (tf *TestFactory) Test() {
	fmt.Print(tf)
}
