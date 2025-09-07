package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println(time.Now(), "Hi from publisher")
		time.Sleep(5 * time.Second)
	}
}
