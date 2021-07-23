package pointersstructsmethods

import "fmt"

func PointersPractice() string {
	// declaringPtrs()
	// dereferencingPtrs()
	passingPtrToFunc()
	return "end of ptrs"
}
func declaringPtrs() {
	b := 255
	var a *int = &b
	fmt.Printf("Type of a is %T\n", a)
	fmt.Println("address of b is", a)

	var c *int
	if c == nil {
		fmt.Println("c is", c)
		c = &b
		fmt.Println("c after initialization is", c)

	}

	size := new(int)
	fmt.Printf("Size value is %d, type is %T, address is %v\n", *size, size, size)
	*size = 85
	fmt.Println("New size value is", *size) //using * to access val is called dereferencing a [ptr]

}

func dereferencingPtrs() {
	b := 255
	a := &b
	fmt.Println("address of b is", a)
	fmt.Println("value of b is", *a)
	*a++
	fmt.Println("new value of b is", b)
}

func change(val *int) {
	*val = 55
}

func returnPtr() *int {
	i := 5
	return &i
}

func passPtrToArr(arr *[3]int) { //bad practice. use slices
	(*arr)[0] = 90
}

func passSliceLOfArr(sls []int) {
	sls[0] = 95
}

func passingPtrToFunc() {
	a := 58
	fmt.Println("value of a before function call is", a)
	b := &a
	change(b)
	fmt.Println("value of a after function call is", a)

	c := returnPtr()
	fmt.Println("Value of c", *c)

	d := [3]int{89, 90, 91}
	passPtrToArr(&d) //bad practice. use slices
	fmt.Println("val of d", d)
	passSliceLOfArr(d[:])
	fmt.Println("val of d", d)

	// b := [...]int{109, 110, 111}
	// p := &b
	// p++    //error.Ptr arithmetic not supported in Go. in other languages, ptr would have moved to next element in arr

}
