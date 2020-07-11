package main

import (
	"fmt"
)

func main() {
	intStream := make(chan int)
	a := cap(intStream)
	fmt.Println("cap", a)
	go func() {
		defer close(intStream)
		for i := 1; i <= 5; i++ {
			intStream <- i
			fmt.Println("put",i)
		}
	}()
	//time.Sleep(5*time.Second)
	for integer := range intStream {
		fmt.Println("get", integer)
	}
}
