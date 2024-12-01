package main

import (
	"errors"
	"fmt"
	"log"
)

type customErr struct {
	Last string
	err  error
}

func (c *customErr) Error() string {
	return fmt.Sprintf("customError %v,%v", c.Last, c.err)
}
func main() {
	cusErr := customErr{
		Last: "Gopher",
		err:  errors.New("something must error"),
	}
	foo(&cusErr)
}

func foo(e error) {
	log.Println(e.Error())
}
