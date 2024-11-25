package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func (p Person) String() string {
	return fmt.Sprintf("\n%v %v %v\n", p.First, p.Last, p.Age)
}

func main() {

	s := `[{"First":"tomson","Last":"xiu","Age":122},{"First":"lakcon","Last":"luis","Age":111}]`
	var p []Person
	err := json.Unmarshal([]byte(s), &p)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p)
}
