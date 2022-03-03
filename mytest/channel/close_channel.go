package main

import "fmt"

func main() {
	ch := make(chan int, 5)
	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch)
	for i := 0; i < 10; i++ {
		t, ok := <-ch
		fmt.Println(t)
		fmt.Println(ok)

	}

}
