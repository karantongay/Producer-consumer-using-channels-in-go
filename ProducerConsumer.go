package main
/* producer-consumer problem in Go */

import ("fmt"
	"time"
)

var complete = make(chan bool)
var items = make(chan int)

func consume () {
	for {
		msg := <-items
		fmt.Println("Consuming", msg)
		time.Sleep(100 * time.Millisecond)
	}
}

func produce () {
	for i := 0; i < 10; i++ {
		fmt.Println("Producing", i)
		items <- i
	}
	complete <- true
}

func main () {
	go consume()
	go produce()
	<- complete
}
