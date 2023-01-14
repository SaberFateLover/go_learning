package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

//场景：在一个高并发的web服务器中，要限制IP的频繁访问。
//现模拟100个IP同时并发访问服务器，每个IP要重复访问1000次。
//每个IP三分钟之内只能访问一次。修改以下代码完成该过程，要求能成功输出 success:100
type Ban struct {
	sync.RWMutex
	visitIPs map[string]time.Time
}

func NewBan() *Ban {
	return &Ban{visitIPs: make(map[string]time.Time)}
}
func (o *Ban) visit(ip string) bool {
	o.Lock()
	defer o.Unlock()
	//ip不存在直接返回true吧
	if t, ok := o.visitIPs[ip]; ok {
		//ip存在就直接返回true? //3分钟以后了，可以让访问了
		now := time.Now()
		if t.Add(time.Minute * 3).Before(now) {
			o.visitIPs[ip] = now
			return false
		}
		return true
	}
	o.visitIPs[ip] = time.Now()
	return false
}
func main() {
	var success int64 = 0
	ban := NewBan()
	group := sync.WaitGroup{}
	group.Add(100)
	for i := 0; i < 1000; i++ {
		for j := 0; j < 100; j++ {
			go func(g *sync.WaitGroup, j int) {
				ip := fmt.Sprintf("192.168.1.%d", j)
				if !ban.visit(ip) {
					//success++
					atomic.AddInt64(&success, 1)
					group.Done()
				}
			}(&group, j)
		}
	}
	group.Wait()
	fmt.Println("success:", success)
}
