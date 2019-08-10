package fibonacci

import (
	"fmt"
	"testing"
)

func TestFibonacci(t *testing.T) {

	a := 1
	b := 1
	fmt.Print(a)
	for i := 0; i < 5; i++ {
		fmt.Print(" ", b)
		temp := a
		a = b
		b = temp + a
	}
}
