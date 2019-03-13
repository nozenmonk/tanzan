package main

import (
	"fmt"
	"context"
)

//objective to learn contexts in golang
func main(){
	rootCtx := context.Background()
	//Rturns a singleton

	a := context.WithValue(rootCtx, "location", "San Francisco")

	b := context.WithValue(a, "Weather", "foggy")

	c := context.WithValue(a, "Weather", "sunny")

	d := context.WithValue(c, "Weather", "rainy")

	fmt.Printf("a %v %v", a.Value("location"), a.Value("Weather"))
	fmt.Printf("\nb %v %v", b.Value("location"), b.Value("Weather"))
	fmt.Printf("\nc %v %v", c.Value("location"),c.Value("Weather"))

	fmt.Printf("\nd %v %v",d.Value("location"), d.Value("Weather"))
}

