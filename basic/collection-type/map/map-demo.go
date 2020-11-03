package main

import "fmt"

func main(){
	var errorType = map[int]string{
		100: "Continue",
		102: "Processing",
		200: "OK",
		201: "Created",
		202: "Accepted",
		203: "Non-Authoritative Information",
	}
	fmt.Println(errorType)
	fmt.Printf("%#v\n", errorType)
}