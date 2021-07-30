package concurrency

import (
	"fmt"
	"time"
)

func SelectBasics() {
	// selectExample()
	// selectPracticalExample()
	// deadlockInSelectExample()		//causes panic
	// deadlockAvoidanceInSelectExample()
	// nilChannelExample()
	// randomSelectionExample()
	emptySelect() //causes panic

}

func server1(ch chan string) {
	time.Sleep(6 * time.Second)
	ch <- "from server1"
}
func server2(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "from server2"

}
func selectExample() {
	output1 := make(chan string)
	output2 := make(chan string)
	go server1(output1)
	go server2(output2)
	select {
	case s1 := <-output1:
		fmt.Println(s1)
	case s2 := <-output2:
		fmt.Println(s2)
	}
}

func process1(ch chan string) {
	time.Sleep(10500 * time.Millisecond)
	ch <- "process successful"
}

func selectPracticalExample() {
	ch := make(chan string)
	go process1(ch)
	for {
		time.Sleep(1000 * time.Millisecond)
		select {
		case v := <-ch:
			fmt.Println("received value: ", v)
			return
		default:
			fmt.Println("no value received")
		}
	}
}

func deadlockInSelectExample() {
	ch := make(chan string)
	select {
	case <-ch:
		fmt.Println("entered case")
	}
}

func deadlockAvoidanceInSelectExample() {
	ch := make(chan string)
	select {
	case <-ch:
	default:
		fmt.Println("default case executed")
	}
}

func nilChannelExample() {
	var ch chan string
	select {
	case v := <-ch:
		fmt.Println("received value", v)
	default:
		fmt.Println("default case executed")
	}
}

func server3(ch chan string) {
	ch <- "from server1"
}
func server4(ch chan string) {
	ch <- "from server2"

}
func randomSelectionExample() {
	output1 := make(chan string)
	output2 := make(chan string)
	go server3(output1)
	go server4(output2)
	time.Sleep(1 * time.Second)
	select {
	case s3 := <-output1:
		fmt.Println(s3)
	case s3 := <-output2:
		fmt.Println(s3)
	}
}

func emptySelect() {
	select {}
}
