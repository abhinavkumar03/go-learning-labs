package main

import (
	"fmt"
	"time"
)

func processNum(messageChan chan int) {
	fmt.Println("processing message", <-messageChan)
}

func processForNum(numChan chan int) {
	for num := range numChan {
		fmt.Println("processing for message", num)

		time.Sleep(1 * time.Second)
	}
}

func sum(result chan int, num1 int, num2 int) {
	numResult := num1 + num2
	result <- numResult
}

// go routine syncronizer
func task(done chan bool) {
	defer func() { done <- true }()
	fmt.Println("process...")
}

func emailSender(emailChan <-chan string, done chan<- bool) {
	defer func() { done <- true }()

	// emailChan <- "hello@gmail.com" // complier time error, cannot send to recieve only channel
	// <-done // complier time error, cannot recieve to send only channel
	for email := range emailChan {
		fmt.Println("sending email to ", email)
		time.Sleep(time.Second)
	}
}

func main() {
	// 1.
	// messageChan := make(chan int)

	// messageChan <- 0 // blocking

	// msg := <-messageChan
	// fmt.Println(msg)

	// go processNum(messageChan)
	// messageChan <- 5

	// time.Sleep(2 * time.Second)

	// 2.
	// numChan := make(chan int)

	// go processForNum(numChan)

	// for {
	// 	numChan <- rand.Intn(100) // infinite loop for channel
	// }

	// 3.
	// result := make(chan int)
	// go sum(result, 4, 5)
	// res := <-result
	// fmt.Println(res)

	// 4.
	// done := make(chan bool)
	// go task(done)
	// <-done // block. the same features done by wait group can be done by channel

	// 5.
	// emailChan := make(chan string, 100) // buffered channel
	// done := make(chan bool)

	// 5.1 -> sending blocking data
	// emailChan <- "1@gmail.com"
	// emailChan <- "2@gmail.com"

	// fmt.Println(<-emailChan)
	// fmt.Println(<-emailChan) // no dead lock

	// 5.2 -> sending via loop
	// go emailSender(emailChan, done)
	// for i := 0; i < 5; i++ {
	// emailChan <- fmt.Sprintf("%d@gmail.com", i)
	// }

	// fmt.Println("done sending.")

	// this is important
	// close(emailChan)
	// <-done

	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() {
		chan1 <- 10
	}()

	go func() {
		chan2 <- "pong"
	}()

	for i := 0; i < 2; i++ {
		select {
		case chan1Val := <-chan1:
			fmt.Println("recieved data from chan 1", chan1Val)

		case chan2Val := <-chan2:
			fmt.Println("recieved data from chan 2", chan2Val)
		}
	}
}
