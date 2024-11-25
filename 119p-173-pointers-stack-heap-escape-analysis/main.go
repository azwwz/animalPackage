package main

import "fmt"

func foo() int {
	a := 1
	return a
}

func main() {
	b := foo()
	fmt.Println(b) // 可以访问 b，因为 b 是 foo 返回的值
	// 无法访问 a，因为 a 在 foo 返回时已经被销毁
}
