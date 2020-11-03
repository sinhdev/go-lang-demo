package main

import (
	"ccdemo/multithread"
)

func main() {
	go multithread.Say("world")
	multithread.Say("hello")
}
