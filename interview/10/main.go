package main

import (
	"fmt"
)

func main1() {
	//这个和java可以不一样， string 的默认值是""
	//一下代码编译直接报错
	//var x string = nil
	//if x == nil {
	//	x = "default"
	//}
	//fmt.Println(x)
}

//这个iota是和位置有关系
const (
	a = iota //0
	b = iota //1

)
const (
	_ = ""   //0
	c = iota //1
	d = iota //2
)

func main2() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}

type query func(string) string

//只执行一个就返回了呀
func exec(name string, vs ...query) string {
	ch := make(chan string)
	//wg:=sync.WaitGroup{}
	//wg.Add(len(vs))
	fn := func(i int) {
		ch <- vs[i](name)
	}
	for i, _ := range vs {
		go fn(i)
	}
	return <-ch
}

func main3() {
	ret := exec("111", func(n string) string {
		return n + "func1"
	}, func(n string) string {
		return n + "func2"
	}, func(n string) string {
		return n + "func3"
	}, func(n string) string {
		return n + "func4"
	})
	fmt.Println(ret)
}

func main4() {
	str1 := []string{"a", "b", "c"}
	str2 := str1[1:]
	//这个是str2 要看清楚，不是str1
	str2[1] = "new"
	fmt.Println(str1) //a b bew
	str2 = append(str2, "z", "x", "y")
	fmt.Println(str1) // a b new
}

type Student struct {
	Name string
}

//指针类型比较的是指针地址，非指针类型比较的是每个属性的值
func main5() {
	fmt.Println(&Student{Name: "menglu"} == &Student{Name: "menglu"}) //false
	fmt.Println(Student{Name: "menglu"} == Student{Name: "menglu"})   //true
}

func main6() {

	//数组只能与相同纬度长度以及类型的其他数组比较，切片之间不能直接比较。。
	fmt.Println([...]string{"1"} == [...]string{"1"})
	//切片之间是不能比较的
	//fmt.Println([]string{"1"} == []string{"1"})
}

type Student2 struct {
	Age int
}

func main() {
	//map 中的value 是不可寻址的
	kv := map[string]*Student2{"menglu": {Age: 21}}
	kv["menglu"].Age = 22
	s := []Student2{{Age: 21}}
	s[0].Age = 22
	fmt.Println(kv, s)
}
