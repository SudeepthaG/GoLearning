package concurrency

import (
	"fmt"
	"time"
)

func GoroutinesBasics() {
	goroutineExample()
	multipleGoroutineExample()
}
func hello() {
	fmt.Println("Hello world goroutine")
}

// Prefix the function or method call with the keyword go and you will have a new Goroutine running concurrently.
func goroutineExample() {
	go hello()
	time.Sleep(1 * time.Second) //without this the control doesnt wait for goroutine to finish and moves to next stmt immediately as The main Goroutine should be running for any other Goroutines to run.
	// If the main Goroutine terminates then the program will be terminated and no other Goroutine will run.
	fmt.Println("main function")
}

func numbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}
func alphabets() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
}

func multipleGoroutineExample() {
	go numbers()
	go alphabets()                      //these two start at almost same time as main goroutine doesnt wait. so the two goroutines are runnon g simultaneously
	time.Sleep(3000 * time.Millisecond) //main goroutine now waits
	fmt.Println("main terminated")
}
