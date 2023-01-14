package main

import (
	"fmt"
	"sync"
	"time"
)

var mu sync.Mutex
var chain string

//为什么panic 是因为 go 中的锁是不可重入的锁
func main1() {
	chain = "main"
	A()
	fmt.Println(chain)
}
func A() {
	mu.Lock()
	defer mu.Unlock()
	chain = chain + " --> A"
	B()
}
func B() {
	chain = chain + " --> B"
	C()
}
func C() {
	mu.Lock()
	defer mu.Unlock()
	chain = chain + " --> C"
}

//A: 不能编译
//B: 输出 main --> A --> B --> C
//C: 输出 main
//D: panic

var mu2 sync.RWMutex
var count int

func main2() {
	go A1()
	time.Sleep(2 * time.Second)
	mu2.Lock()
	defer mu2.Unlock()
	count++
	fmt.Println(count)
}
func A1() {
	mu2.RLock()
	defer mu2.RUnlock()
	B1()
}
func B1() {
	time.Sleep(5 * time.Second)
	C1()
}
func C1() {
	mu2.RLock()
	defer mu2.RUnlock()
}

//A: 不能编译
//B: 输出 1
//C: 程序hang住
//D: panic

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		time.Sleep(time.Millisecond)
		wg.Done()
		//wg.Wait 执行完成之后再添加就直接panic了
		wg.Add(1)
	}()
	wg.Wait()
}

//A: 不能编译
//B: 无输出，正常退出
//C: 程序hang住
//D: panic

type Once struct {
	m    sync.Mutex
	done uint32
}

func (o *Once) Do(f func()) {
	if o.done == 1 {
		return
	}
	o.m.Lock()
	defer o.m.Unlock()
	if o.done == 0 {
		o.done = 1
		f()
	}
}

//A: 不能编译
//B: 可以编译，正确实现了单例
//C: 可以编译，有并发问题，f函数可能会被执行多次
//D: 可以编译，但是程序运行会panic
