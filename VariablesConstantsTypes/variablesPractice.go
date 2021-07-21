package variablesconstantstypes

import (
	"fmt"
	"unsafe"
)

func VariablesFunc() string {

	// types of declaring and initializing variables

	fmt.Println("Single variables")
	var age int
	var stud = 3
	age = 10
	var age2 = 20
	fmt.Println(age, " ", age2, " ", stud)

	fmt.Println("Multiple variables of same type")
	var ht1, w1 int
	ht1, w1 = 10, 20
	var ht2, w2 = 10, 20
	ht3, w3 := 10, 20
	fmt.Println(ht1, " ", w1, " ", ht2, " ", w2, " ", ht3, " ", w3, " ")
	ht3, w4 := 20, 30
	// ht3,w3 =10,20 gives error as there are no new variables
	fmt.Println(ht3, " ", w4, " ")

	fmt.Println("Multiple variables of different type")
	name, no := "GS", 23
	fmt.Println(name, " ", no)

	var (
		x   = 10
		y   = "y"
		z   = 10.3
		xyz int
	)
	fmt.Println(x, y, z, " ", xyz)

	fmt.Println("displaying variable types and sizes")
	var a = 10
	fmt.Println(a)
	var str = "adc"
	fmt.Printf("type is %T\n", str) //doesnt print type if you use println/print. u should use printf
	fmt.Printf("size is %d", unsafe.Sizeof(str))
	anInt := true
	fmt.Printf("\n%T\n", anInt)
	fmt.Printf("size is %d", unsafe.Sizeof(anInt))
	c := 23
	fmt.Println(c, "%T\n", c, "size is %d", unsafe.Sizeof(c))
	d := 10.3
	// printf will not allow you to concat  numbers with strings and will not let u use more than 1 argument
	// fmt.Printf(d, "%T\n", d, "size is %d", unsafe.Sizeof(c))
	// fmt.Printf("%T\n", d, "size is ", unsafe.Sizeof(c))
	fmt.Printf("%T\n", d)
	fmt.Printf("size is %d", unsafe.Sizeof(c))

	fmt.Println("complexnumbers")
	c1 := complex(1, 2)
	c2 := 3 + 4i
	fmt.Print(c1)
	fmt.Print(c2)

	fmt.Println("Type conversion")
	i := 1
	j := 2.4
	// k:=i+j   //gives error.Go doesnt do auto type cast
	k := i + int(j)
	l := float64(i) + j
	println(i)
	println(j)
	println(k)
	println(l)
	return "End of Variables function"

}
