package variablesconstantstypes

import (
	"fmt"
)

func ConstantsFunc() string {
	const str = "abcd"
	fmt.Printf("type is %T", str)
	const (
		name    = "John"
		age     = 50
		country = "Canada"
	)
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(country)

	// var a = math.Sqrt(4) //allowed
	// const b = math.Sqrt(4) //not allowed

	const n = "Sam"
	var namen = n
	fmt.Printf("type %T value %v \n", namen, namen)

	var defaultName = "Sam" //allowed
	type myString string
	var customName myString = "Sam" //allowed
	// customName = defaultName //not allowed as types are different
	fmt.Printf("\ntype %T value %T\n", defaultName, customName)

	const trueConst = true
	type myBool bool
	var defaultBool = trueConst       //allowed
	var customBool myBool = trueConst //allowed
	// defaultBool = customBool //not allowed
	println(defaultBool)
	println(customBool)

	// default type of these kinds of constants can be thought of as being generated on the fly depending on the context.
	const a = 5
	var intVar int = a
	var int32Var int32 = a
	var float64Var float64 = a
	var complex64Var complex64 = a
	fmt.Println("intVar", intVar, "\nint32Var", int32Var, "\nfloat64Var", float64Var, "\ncomplex64Var", complex64Var)

	// 5.9/8 is allowed as both are numeric constants.
	var div = 5.9 / 8
	fmt.Printf("a's type is %T and value is %v", div, div)

	return "\nEnd of contants Function"

}
