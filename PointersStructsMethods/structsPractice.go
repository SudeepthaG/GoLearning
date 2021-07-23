package pointersstructsmethods

import "fmt"

type Spec struct { //exported struct. can be accessed in other packages. It is a good practice to write all  globally accessible structs in a different module->file
	Maker string //exported field
	Price int    //exported field
	model string //unexported field - can't be accessed outside

}

func StructsPractice() string {

	compareStructs()
	anonymous_nested_promoted_fields()
	return "end of structs"
}

type Person struct {
	string
	int               //since these are anonymous fields we can't redeclare another string or int
	address   Address //this is nested struct
	P_Address         // its fields are now promoted

}
type Address struct {
	city  string
	state string
}

type P_Address struct {
	city  string
	state string
}

func anonymous_nested_promoted_fields() {
	// Even though anonymous fields do not have an explicit name, by default the name of an anonymous field is the name of its type.
	// Fields that belong to an anonymous struct field in a struct are called promoted fields since they can be accessed as if they belong to the struct which holds the anonymous struct field.
	p := Person{
		string: "naveen",
		int:    50,
		address: Address{
			city:  "Chicago",
			state: "Illinois",
		},
		P_Address: P_Address{
			city:  "Chic",
			state: "Illi",
		},
	}
	fmt.Println(p.string)
	fmt.Println(p.int)
	fmt.Println("City:", p.address.city)
	fmt.Println("State:", p.address.state)
	fmt.Println("City:", p.city)   //city is promoted field
	fmt.Println("State:", p.state) //state is promoted field

}

type name struct {
	firstName string
	lastName  string
}

type image struct {
	data map[int]int
}

func compareStructs() {

	// Structs are value types and are comparable if each of their fields are comparable.
	// Two struct variables are considered equal if their corresponding fields are equal.
	name1 := name{
		firstName: "Steve",
		lastName:  "Jobs",
	}
	name2 := name{
		firstName: "Steve",
		lastName:  "Jobs",
	}
	if name1 == name2 {
		fmt.Println("name1 and name2 are equal")
	} else {
		fmt.Println("name1 and name2 are not equal")
	}

	name3 := name{
		"Steve", //we can declare directly also like this,but it is not preferred as in this case, it assigns the values in the order they are declared and it is difficult to understand which input is for which field
		"Jobs",
	}
	name4 := name{
		firstName: "Steve",
	}

	if name3 == name4 {
		fmt.Println("name3 and name4 are equal")
	} else {
		fmt.Println("name3 and name4 are not equal")
	}

	// image struct type contains a field data which is of type map. maps are not comparable, hence image1 and image2 cannot be compared.
	// image1 := image{
	//     data: map[int]int{
	//         0: 155,
	//     }}
	// image2 := image{
	//     data: map[int]int{
	//         0: 155,
	//     }}
	// if image1 == image2 { //error
	//     fmt.Println("image1 and image2 are equal")
	// }

}
