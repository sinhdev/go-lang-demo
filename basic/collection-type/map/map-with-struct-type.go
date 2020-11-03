package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["VTCA TT"] = Vertex{20.9922311, 105.8601709}
	m["VTCA LĐH"] = Vertex{10.7608141, 106.6569949}
	fmt.Println(m["VTCA LĐH"])
}
