package main

import "sync"  	//for mutex lock, unlock and waitgroups
import "fmt"
import "time"	//for sleep

//starting main function
func main(){

	type value struct{	//value has both an int and a Mutex
		mu sync.Mutex
		value int
	}
	
	var wg sync.WaitGroup		//wg is a WaitGroup

	printSum := func(v1, v2 *value){ 	//v1 v2 are passed by reference
		defer wg.Done()			//Done() reduces Add count by 1
		v1.mu.Lock()			//Lock added for v1
		defer v1.mu.Unlock()		//v1 UnLock defered
		
		time.Sleep(2 * time.Second)	//sleep for 2 seconds

		v2.mu.Lock()			//v2 Lock
		defer v2.mu.Unlock()		//v2 Unlock defered

		fmt.Printf("sum=%v \n", v1.value + v2.value)
	}
	
	var a, b value
	wg.Add(2)			//waitgroup waits for 2 Dones
	go printSum(&a, &b)
	go printSum(&b, &a)
	wg.Wait()
}
	
