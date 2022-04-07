package main

import (
	"fmt"
)

func main() {
	pointer()
	pointer_nil()
	pointer_new()
	pointer_dereferencing()
	pointer_dereferencing1()
	function_passing_pointer()
	function_returning_pointer()
	function_passing_array()
	function_passing_slices()
	pointer_arithmetic()

}

//*T is the type of the pointer variable which points to a value of type T.
func pointer() {
	b := 255
	var a *int = &b
	fmt.Printf("Type of a is %T\n", a)
	fmt.Println("address of b is", a)
	fmt.Println("a = ", *a)
}

func pointer_nil() {
	a := 25
	var b *int
	if b == nil {
		fmt.Println("b is", b)
		b = &a
		fmt.Println("b after initialization is", b)
		fmt.Println("b = ", *b)
	}
}

//Creating pointers using the new function
func pointer_new() {
	size := new(int)
	fmt.Printf("Size value is %d, type is %T, address is %v\n", *size, size, size)
	*size = 85
	fmt.Println("New size value is", *size)
}

func pointer_dereferencing() {
	b := 150
	a := &b
	fmt.Println("address of b is", a)
	fmt.Println("value of b is", *a)
}

func pointer_dereferencing1() {
	b := 255
	a := &b
	fmt.Println()
	fmt.Println("address of b is", a)
	fmt.Println("value of b is", *a)
	*a++
	fmt.Println("new value of b is", b)
}

func pointer_change(c *int) {
	*c = 55
}
func function_passing_pointer() {
	a := 58
	fmt.Println("a = ", a)
	b := &a
	pointer_change(b)
	fmt.Println("a = ", a)
}

func pointer_hello() *int {
	i := 15
	return &i
}
func function_returning_pointer() {
	fmt.Println()
	d := pointer_hello()
	fmt.Println("d = ", *d)
}

/*Do not pass a pointer to an array as a argument to a function. Use slice instead.*/
func pointer_modify_array(arr *[3]int) {
	(*arr)[0] = 90
	//arr[0] = 90	//also working
}

func function_passing_array() {
	a := [3]int{89, 90, 91}
	pointer_modify_array(&a)
	fmt.Println("by array", a)
}

/*Although this way of passing a pointer to an array as a argument to a function and making modification is working but,
it is not the idiomatic way. In Go slices is used to achieve this.
*/
func pointer_modify_slices(sls []int) {
	sls[0] = 90
}
func function_passing_slices() {
	a := [3]int{89, 90, 91}
	pointer_modify_slices(a[:])
	fmt.Println("by slice", a)
}

/*Go does not support pointer arithmetic which is present in other languages like C and C++.*/
func pointer_arithmetic() {
	b := [...]int{109, 110, 111}
	p := &b

	//p++	//arithmetic calcultion is not allowed over pointer

	fmt.Println(b)
	fmt.Println(p)
	fmt.Println(*p)
}

//https://go.dev/play/p/WtLJs6QN5IV