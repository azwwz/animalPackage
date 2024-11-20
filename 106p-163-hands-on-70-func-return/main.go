package main

import "fmt"

func ReturnFunc() func() int{
	return func ()int{
		return 42
	}
}
func main(){
	var f func() int = ReturnFunc()
	fmt.Println(f())
}

