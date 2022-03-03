package main

import (
	"fmt"
	"time"
)

func main() {
	// messages := make(chan int, 10)
	// done := make(chan bool)

	// defer close(message)
	// // consumer
	// go func() {
	//     ticker := time.NewTicker(1 * time.Second)
	//     for _ = range ticker.C {
	//         select {
	//         case <-done:
	//             fmt.Println("child process interrupt...")
	//             return
	//         default:
	//             fmt.Printf("send message: %d\n", <-messages)
	//         }
	//     }
	// }()

	// producer
	// for i := 0; i < 10; i++ {
	//     messages <- i
	// }
	// time.Sleep(5 * time.Second)
	// close(done)
	// time.Sleep(1 * time.Second)

	ticker := time.NewTicker(2 * time.Second)
	// for _=range ticker.C {
	// 	fmt.Println("send message",<-messages)
	// }
	go func() {
		for {
			//从定时器中获取数据
			t := <-ticker.C
			fmt.Println("当前时间为:", t)

		}
	}()
	for {
		time.Sleep(time.Second * 2)
	}
}
