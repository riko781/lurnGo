package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	/**************/
	//code à excuter
	/*************/
	finish := time.Since(start)
	fmt.Println(finish)
}
