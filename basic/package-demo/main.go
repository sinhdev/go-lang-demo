package main

import (
	"fmt"
	"packagedemo/greetings"
)

func main() {
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
