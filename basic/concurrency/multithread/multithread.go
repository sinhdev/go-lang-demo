package multithread

import (
	"fmt"
	"time"
)

//function say
func Say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}
