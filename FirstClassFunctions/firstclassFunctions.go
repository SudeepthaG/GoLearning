package firstclassfunctions

import "fmt"

func FistClassFunctionsPractice() {
	anonymousFunctionsExample()
	userDefinedFunctionTypes()
	functionAsArgument()
	returningFunctionsExample()
	closureExample()
	closureExample2()
	closurePracticalExample()
	mappingFunction()
}

func anonymousFunctionsExample() {
	a := func() {
		fmt.Println("hello world first class function")
	}
	a()
	fmt.Printf("%T", a)

	func() {
		fmt.Println("\nhello world first class function")
	}()

	func(n string) {
		fmt.Println("Welcome", n)
	}("Gophers")
}

type add func(a int, b int) int

func userDefinedFunctionTypes() {
	var a add = func(a int, b int) int {
		return a + b
	}
	s := a(5, 6)
	fmt.Println("Sum", s)
}

// Higher-order functions: a function which does at least one:-takes one or more functions as arguments
// returns a function as its result
func simple(a func(a, b int) int) {
	fmt.Println("sum")
	fmt.Println(a(60, 7))
}

func functionAsArgument() {
	f := func(a, b int) int {
		fmt.Println("sum2")
		return a + b
	}
	simple(f)
}

func simple2() func(a, b int) int {
	f := func(a, b int) int {
		return a + b
	}
	return f
}

func returningFunctionsExample() {
	s := simple2()
	fmt.Println(s(50, 7))
}

// Closures are anonymous functions that access the variables defined outside the body of the function.
func closureExample() {
	a := 5
	func() {
		fmt.Println("a =", a)
	}()
}

func appendStr() func(string) string {
	t := "Hello"
	c := func(b string) string {
		t = t + " " + b
		return t
	}
	return c

}
func closureExample2() {
	a := appendStr()
	b := appendStr()
	fmt.Println(a("World"))
	fmt.Println(b("everyone"))
	fmt.Println(a("gopher"))
	fmt.Println(b("!"))
}

type student struct {
	firstName string
	lastName  string
	grade     string
	country   string
}

func filter(s []student, f func(student) bool) []student {
	var r []student
	for _, v := range s {
		if f(v) == true {
			r = append(r, v)
		}
	}
	return r
}

func closurePracticalExample() {
	s1 := student{
		firstName: "Naveen",
		lastName:  "Ramanathan",
		grade:     "A",
		country:   "India",
	}
	s2 := student{
		firstName: "Samuel",
		lastName:  "Johnson",
		grade:     "B",
		country:   "USA",
	}
	s := []student{s1, s2}
	f := filter(s, func(s student) bool {
		if s.grade == "B" {
			return true
		}
		return false
	})
	fmt.Println(f)
	c := filter(s, func(s student) bool {
		if s.country == "India" {
			return true
		}
		return false
	})
	fmt.Println(c)
}

func iMap(s []int, f func(int) int) []int {
	var r []int
	for _, v := range s {
		r = append(r, f(v))
	}
	return r
}

func mappingFunction() {
	a := []int{5, 6, 7, 8, 9}
	r := iMap(a, func(n int) int {
		return n * 5
	})
	fmt.Println(r)
}
