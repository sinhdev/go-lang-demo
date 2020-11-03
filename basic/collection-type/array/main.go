package main

import "fmt"

func main(){
	var a1 [2]string
	a1[0] = "Hello"
	a1[1] = "World"
	fmt.Printf("%s %s\n", a1[0], a1[1]);
	fmt.Println(a1)
	
	var a2 = [2]string{"Xin", "chào"}
	fmt.Println(a2)
}