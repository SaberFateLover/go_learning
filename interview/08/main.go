package main

import (
	"fmt"
	"time"
)

//time.Tick
func main() {
	go func() {
		// 1 在这里需要你写算法
		// 2 要求每秒钟调用一次proc函数
		// 3 要求程序不能退出

		t := time.NewTicker(time.Second * 1) //每个Duration 产生一个chan
		for {
			select {
			case <-t.C:
				go func() {
					//这个要放在函数里面，原来是这样
					defer func() {
						if err := recover(); err != nil {
							fmt.Println("hah", err)
						}
					}()
					proc()
				}()
			}
		}
	}()

	select {}
}

func proc() {
	panic("ok")
	//后面不走了
	fmt.Println(".....")
}
