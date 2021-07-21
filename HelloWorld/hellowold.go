package helloworld

import "fmt"

// function name should always start with capital letter

func Hello() string { //If you are not returning anything , u will write func Hello(){ ... }
	fmt.Println("This is going to print our standard first program")
	return "Hello, World"
}
