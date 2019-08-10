package main

import (
	"fmt"
)

func main() {
	var b = true
	if b {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
	expr := "expr"
	switch {
	case 2 < len(expr):
		fmt.Println(expr, "length =", len(expr), ">2")
		fallthrough
	case 5 < len(expr):
		fmt.Println(expr, "length =", len(expr), ">5")
		fallthrough
	case false:
		fmt.Println("break")
		if true {
			break
		}
		fmt.Println("???")
		fallthrough
	case true:
		fmt.Println("oh my ass")
	default:
		fmt.Println("default")
	}

	var x interface{}
	switch i := x.(type) {
	case nil:
		fmt.Println(" x 的类型 :%T", i)
	case int:
		fmt.Println("x 是 int 型")
	case float64:
		fmt.Println("x 是 float64 型")
	case func(int) float64:
		fmt.Println("x 是 func(int) 型")
	case bool, string:
		fmt.Println("x 是 bool 或 string 型")
	default:
		fmt.Println("未知型")
	}

	/**
	SELECT
	*/
	var c1, c2, c3 chan int
	var i1, i2 int
	select {
	case i1 = <-c1:
		fmt.Println("received ", i1, " from c1\n")
	case c2 <- i2:
		fmt.Println("sent ", i2, " to c2\n")
	case i3, ok := (<-c3): // same as: i3, ok := <-c3
		if ok {
			fmt.Println("received ", i3, " from c3\n")
		} else {
			fmt.Println("c3 is closed\n")
		}
	default:
		fmt.Println("no communication\n")
	}
	/**
	  TODO chan interface
	*/
}
