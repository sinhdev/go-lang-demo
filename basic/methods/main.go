package main

import (
	"fmt"
	"methods/method1"
)

func main() {
	u := method1.User{"VTC", "Academy"}
	fmt.Println(u.Greeting())
}
