package concurrency

import (
	"fmt"
	"time"
)

func ChannelsBasics() {
	// channelExample()
	// channelWriteReceiveExample()
	// channelWithSleepExample()
	// twoChannelsExample()
	// deadlockExample() //this causes panic
	// unidirectionalChannelExample()
	// closeAndForLoopsExample()
	// closeAndForLoopsShortExample()
	readableTwoChannelsExample()

}

func channelExample() {

	var a chan int
	if a == nil {
		fmt.Println("channel a is nil, going to define it")
		a = make(chan int)
		fmt.Printf("Type of a is %T", a)
	}
	// a := make(chan int)  //or this can be done
}

func hello2(done chan bool) {
	fmt.Println("Hello world goroutine")
	done <- true
}

func channelWriteReceiveExample() {
	done := make(chan bool)
	go hello2(done)
	<-done //try commmenting this line and see.without this blocking is not done cause channels are blocking by default whenw e read or write dta
	fmt.Println("main function")
}

func hello3(done chan bool) {
	fmt.Println("hello go routine is going to sleep")
	time.Sleep(4 * time.Second)
	fmt.Println("hello go routine awake and going to write to done")
	done <- true
}

func channelWithSleepExample() {
	done := make(chan bool)
	fmt.Println("Main going to call hello go goroutine")
	go hello3(done)
	<-done
	fmt.Println("Main received data")
}

func calcSquares(number int, squareop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit
		number /= 10
	}
	squareop <- sum
}

func calcCubes(number int, cubeop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit * digit
		number /= 10
	}
	cubeop <- sum
}

func twoChannelsExample() {
	number := 589
	sqrch := make(chan int)
	cubech := make(chan int)
	go calcSquares(number, sqrch)
	go calcCubes(number, cubech)
	squares, cubes := <-sqrch, <-cubech
	fmt.Println("Final output", squares+cubes)
}

// if a Goroutine is waiting to receive data from a channel, then some other Goroutine is expected to write data on that channel, else the program will panic.
func deadlockExample() {
	ch := make(chan int)
	ch <- 5
}

func sendData(sendch chan<- int) {
	sendch <- 10
}

func unidirectionalChannelExample() {
	// sendch := make(chan<- int)
	// go sendData(sendch)
	// fmt.Println(<-sendch)	//error we cant receive from a send pnly channel
	chnl := make(chan int)
	go sendData(chnl)
	fmt.Println(<-chnl) //here chnl is bidirectional here nad send only in sendData routine

}

func producer(chnl chan int) {
	for i := 0; i < 10; i++ {
		fmt.Println("in producer loop")
		chnl <- i
	}
	close(chnl)
}
func closeAndForLoopsExample() {
	ch := make(chan int)
	go producer(ch)
	for {
		fmt.Println("in main loop")
		v, ok := <-ch
		if ok == false {
			break
		}
		fmt.Println("Received ", v, ok)
	}
}

func closeAndForLoopsShortExample() {
	ch := make(chan int)
	go producer(ch)
	for v := range ch {
		fmt.Println("Received ", v)
	}
}

func digits(number int, dchnl chan int) {
	for number != 0 {
		digit := number % 10
		dchnl <- digit
		number /= 10
	}
	close(dchnl)
}
func calcSquares2(number int, squareop chan int) {
	sum := 0
	dch := make(chan int)
	go digits(number, dch)
	for digit := range dch {
		sum += digit * digit
	}
	squareop <- sum
}

func calcCubes2(number int, cubeop chan int) {
	sum := 0
	dch := make(chan int)
	go digits(number, dch)
	for digit := range dch {
		sum += digit * digit * digit
	}
	cubeop <- sum
}

func readableTwoChannelsExample() {
	number := 589
	sqrch := make(chan int)
	cubech := make(chan int)
	go calcSquares2(number, sqrch)
	go calcCubes2(number, cubech)
	squares, cubes := <-sqrch, <-cubech
	fmt.Println("Final output", squares+cubes)
}
