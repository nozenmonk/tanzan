package main

import (
	"sync"
	"fmt"
)

func main() {
	var memoryAccess sync.Mutex
	var value int
	go func() {
		memoryAccess.Lock()
		value++
		memoryAccess.Unlock()
	}()
	
	memoryAccess.Lock()
	if value == 0 {
		fmt.Println("the value is", value)
	} else {
		fmt.Println("the value is", value)
	}
	memoryAccess.Unlock()
}
 
	

