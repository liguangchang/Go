package main

import (
	"fmt"
	"os"
)

/**
程序入口为package 为 main的不支持任何返回值main()函数
os.Exit返回状态
*/
func main() {
	println("-----")
	fmt.Println(os.Args)
	if len(os.Args) > 1 {
		fmt.Print(os.Args[1])
		fmt.Println(os.Args[1])
		print(os.Args[1])
		fmt.Printf("hello go")
	}
}
