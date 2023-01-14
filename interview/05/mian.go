package main

import (
	"math/rand"
	"time"
)

//写代码实现两个 goroutine，其中一个产生随机数并写入到 go channel 中，
//另外一个从 channel 中读取数字并打印到标准输出。最终输出五个随机数。

//让主线程暂停，可以使用waitGroup
func main() {

	chans := make(chan int)

	stop := make(chan struct{})

	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(time.Second * 1)
			ran := rand.Int()
			chans <- ran
		}
		stop <- struct{}{}
	}()

	//for 和 select 在一起，break 只能break select 如果想直接退出，只能loop: 这种形式。
	go func() {
	loop:
		for {
			select {
			case item := <-chans:
				println(item)
			case <-stop:
				println("结束了")
				break loop
			}
		}
		println("循环结束了！")
	}()

	time.Sleep(time.Second * 7)
}
