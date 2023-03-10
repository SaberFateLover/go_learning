package main

import (
	"fmt"
)

func IndexArray() {
	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8}

	for i := range a {
		a[3] = 100
		if i == 3 {
			fmt.Println("IndexArray", i, a[i])
		}
	}
}

func IndexValueArray() {
	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8}

	//这个是副本，这是啥原理
	for i, v := range a {
		a[3] = 100
		if i == 3 {
			fmt.Println("IndexValueArray", i, v)
		}
	}
}

func IndexValueArrayPtr() {
	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8}

	for i, v := range &a {
		a[3] = 100
		if i == 3 {
			fmt.Println("IndexValueArrayPtr", i, v)
		}
	}
}

func main() {
	IndexArray()         //3 100
	IndexValueArray()    //3 4
	IndexValueArrayPtr() //3 100
}
