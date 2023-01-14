//https://github.com/lifei6671/interview-go/blob/master/question/q007.md
package main

import (
	"fmt"
	"reflect"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

//func main() {
//
//}

type Param map[string]interface{}

type Show struct {
	Param
}

//main1 这个不是主函数
func main1() {
	//创建的是个指针Param 本质是个map ,不能用new 初始化
	//s := new(Show)
	//s.Param["RMB"] = 10000

	s := &student{}
	zhoujielun(s)
}

//type student struct {
//	Name string
//}

func zhoujielun(v interface{}) {
	switch msg := v.(type) {
	case *student, student:
		fmt.Printf("%v", msg)
	}
}

//私有属性不应该加tag的
//转码可能有问题
/*
type People struct {
	//Struct field 'name' has 'json' tag but is not exported
	name string `json:"name"`
}

func main1() {
	js := `{
		"name":"11"
	}`
	var p People
	err := json.Unmarshal([]byte(js), &p)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	fmt.Println("people: ", p)
}
*/

type People struct {
	Name string
}

//Placeholder argument causes a recursive call to the 'String' method (%v)
func (p *People) String() string {
	return fmt.Sprintf("print: %v", p)
}

//fatal error: stack overflow
func main2() {
	p := &People{}
	p.String()
}

//panic: send on closed channel
func main3() {
	ch := make(chan int, 1000)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()
	go func() {
		for {
			a, ok := <-ch
			if !ok {
				fmt.Println("close")
				return
			}
			fmt.Println("a: ", a)
		}
	}()
	close(ch)
	fmt.Println("ok")
	time.Sleep(time.Second * 100)
}

var value int32

func SetValue(delta int32) {
	//CompareAndSwapInt32 这个不需要加for是对的
	for {
		v := value
		if atomic.CompareAndSwapInt32(&value, v, v+delta) {
			break
		}
	}
}

type Project struct{}

func (p *Project) deferError() {
	if err := recover(); err != nil {
		fmt.Println("recover: ", err)
	}
}

func (p *Project) exec(msgchan chan interface{}) {
	for msg := range msgchan {
		//传过来的是string ,用 int 判断会抛异常吧
		m := msg.(int)
		fmt.Println("msg: ", m)
	}
}

//Possible resource leak, 'defer' is called in the 'for' loop
func (p *Project) run(msgchan chan interface{}) {
	defer p.deferError()
	for {
		go p.exec(msgchan)
		time.Sleep(time.Second * 2)
	}
}

func (p *Project) Main() {
	a := make(chan interface{}, 100)
	go p.run(a)
	go func() {
		for {
			a <- "1"
			time.Sleep(time.Second)
		}
	}()
	//1<<63-1
	time.Sleep(time.Second * 10000) //太大了
}

func main4() {
	p := new(Project)
	p.Main()
}

func main5() {
	abc := make(chan int, 1000)
	for i := 0; i < 10; i++ {
		abc <- i
	}
	go func() {
		for a := range abc {
			fmt.Println("a: ", a)
		}
	}()
	close(abc)
	fmt.Println("close")
	time.Sleep(time.Second * 100)
}

//type Student struct {
//	name string
//}

func main6() {
	//m := map[string]Student{"people": {"zhoujielun"}}
	//便衣就会报错
	//m["people"].name = "wuyanzu"
	//map的value本身是不可寻址的，因为map中的值会在内存中移动，并且旧的指针地址在map改变时会变得无效。
	//故如果需要修改map值，可以将map中的非指针类型value，修改为指针类型，比如使用map[string]*Student
}

type query func(string) string

func exec(name string, vs ...query) string {
	ch := make(chan string)
	fn := func(i int) {
		ch <- vs[i](name)
	}
	for i, _ := range vs {
		go fn(i)
	}

	//这样写应该会返回一个值吧
	return <-ch
}

func main8() {
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

func main9() {
	//byte 其实是uint8 ,无符号的8为最大才255，所以i <=255 为true
	//var i byte
	go func() {
		//for i = 0; i <= 255; i++ {
		//}
		for i := 0; i < 10; i++ {
			println(i)
			//time.Sleep(time.Second*2)
		}
	}()
	fmt.Println("Dropping mic")
	// Yield execution to force executing other goroutines
	runtime.Gosched()
	runtime.GC()
	fmt.Println("Done")

	//正在被执行的 goroutine 发生以下情况时让出当前 goroutine 的执行权，并调度后面的 goroutine 执行：
	//IO 操作
	//Channel 阻塞
	//system call
	//运行较长时间
	//如果一个 goroutine 执行时间太长，scheduler 会在其 G 对象上打上一个标志（ preempt）//抢占的意思，
	//当这个 goroutine 内部发生函数调用的时候，会先主动检查这个标志，如果为 true 则会让出执行权。
	//main 函数里启动的 goroutine 其实是一个没有 IO 阻塞、没有 Channel 阻塞、没有 system call、没有函数调用的死循环。
	//也就是，它无法主动让出自己的执行权，即使已经执行很长时间，scheduler 已经标志了 preempt
	//而 golang 的 GC 动作是需要所有正在运行 goroutine 都停止后进行的。因此，程序会卡在 runtime.GC() 等待所有协程退出。
}

type student struct {
	Name string
	Age  int
}

func pase_student() {

	m := make(map[string]*student)

	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}

	//
	for _, stu := range stus {
		m[stu.Name] = &stu
	}
	fmt.Printf("%v", stus)

	//golang的for ... range语法中，stu变量会被复用，
	//每次循环会将集合中的值复制给这个变量，
	//因此，会导致最后m中的map中储存的都是stus最后一个student的值。

}

func main10() {
	pase_student()
}

func main11() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)

	//Loop variable 'i' captured by the func literal
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("i: ", i) //10 10 10 10 10 10 .....
			wg.Done()
		}()
	}

	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("i: ", i) //1,2 ,3,4 ,5 ,6 . . . .
			wg.Done()
		}(i)
	}

	wg.Wait()
}

// 输出结果为showA、showB。
// golang 语言中没有继承概念，
// 只有组合，也没有虚方法，更没有重载。因此，*Teacher 的 ShowB
// 不会覆写被组合的 People 的方法。
type People2 struct{}

func (p *People2) ShowA() {
	fmt.Println("showA")
	//这里和java不同，调用的是People 的showB方法
	p.ShowB()
}
func (p *People2) ShowB() {
	fmt.Println("showB")
	//这里好像没办法调用 teacher 的 ShowB方法
}

type Teacher struct {
	People2
}

//应该是输出  showB
func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

func main12() {
	t := Teacher{}
	t.ShowB()

	t.ShowA()
}

//下面代码会触发异常吗？请详细说明
func main13() {
	runtime.GOMAXPROCS(1)

	//
	int_chan := make(chan int, 1)

	//
	string_chan := make(chan string, 1)

	int_chan <- 1
	string_chan <- "hello"

	//结果是随机执行。golang 在多个case 可读的时候会公平的选中一个执行
	//select 中的case 是随机的
	select {

	case value := <-int_chan:
		//
		fmt.Println(value)
	case value := <-string_chan:

		//
		panic(value)
	}

}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

//以后少吵架
//没事就别吵架了
//想想我就头大

//defer 在定义的时候会计算好调用函数的参数，所以会优先输出10、20 两个参数。然后根据定义的顺序倒序执行。
func main14() {
	a := 1
	b := 2

	//传入的是啥值就是啥值
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1
}

func main15() {
	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s) //0 , 0, 0, 0, 0, 1, 2, 3
}

type UserAges struct {
	ages map[string]int
	//读写🔒
	sync.RWMutex
}

func (ua *UserAges) Add(name string, age int) {
	ua.Lock()
	defer ua.Unlock()
	//这个是不是要用指针
	ua.ages[name] = age
}

//fatal error: concurrent map read and map write
//map 不是线程安全的，所以读的时候也要加锁
func (ua *UserAges) Get(name string) int {
	ua.Lock()
	defer ua.Unlock()
	if age, ok := ua.ages[name]; ok {
		return age
	}
	return -1
}

type People3 interface {
	Speak(string) string
}

type Student struct{}

func (stu *Student) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}

func main16() {
	//这样是可以的
	s := Student{}

	var p People3 = &Student{}
	think := "bitch"
	fmt.Println(s.Speak(think))
	fmt.Println(p.Speak(think))
}

//编译失败，值类型 Student{} 未实现接口People的方法，不能定义为 People类型。
//在 golang 语言中，Student 和 *Student 是两种类型，第一个是表示 Student 本身，第二个是指向 Student 的指针。
//但是指针类型的receiver 是可以接受对象类型。

type People4 interface {
	Show()
}

type Student4 struct{}

func (stu *Student4) Show() {

}

func live() People4 {
	var stu *Student4
	return stu
}

func main() {
	l := live()
	//如果想判断接口的值是不是nil, 那我们就直接用反射就可以了
	if reflect.ValueOf(l).IsNil() {
		fmt.Println("AAAAAAA")
	} else {
		fmt.Println("BBBBBBB")
	}

	//这一点和java 不同，这个空接口的值为nil 但是类型不是nil  所以 l!=nil

	//跟上一题一样，不同的是*Student 的定义后本身没有初始化值，
	//所以 *Student 是 nil的，但是*Student 实现了 People 接口，接口不为 nil
}
