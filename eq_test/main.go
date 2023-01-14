package main

import "fmt"

func main() {
	//f1()
	//f2

	a := make([]int, 0, 0)

	//分配新的切片
	b := append(a, 1)

	fmt.Printf("%+v\n", b)
	//也是在a的基础上分配新的切片
	a = append(a, 2)

	fmt.Printf("%+v\n", a)

}

func f1() {
	var a, b struct{}
	//print(&a, "\n", &b, "\n") // Prints same address
	fmt.Println(&a == &b) // Comparison returns false
}

func f2() {
	var a, b struct{}
	fmt.Printf("%p\n%p\n", &a, &b) // Again, same address
	fmt.Println(&a == &b)          // ...but the comparison returns true
}
