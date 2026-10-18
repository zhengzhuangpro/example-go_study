package main

import (
	"fmt"
	"sync"
)

func main() {
	var maps = sync.Map{}

	go func() {
		for {
			maps.Store(2, "b")
		}
	}()

	go func() {
		for {
			val, ok := maps.Load(2)
			fmt.Println(val, ok)
		}
	}()

	select {}
}
