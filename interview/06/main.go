package main

import (
	"fmt"
	"sync"
	"time"
)

type sp interface {
	Out(key string, val interface{})                  //存入key /val，如果该key读取的goroutine挂起，则唤醒。此方法不会阻塞，时刻都可以立即执行并返回
	Rd(key string, timeout time.Duration) interface{} //读取一个key，如果key不存在阻塞，等待key存在或者超时
}

type SyncMap struct {
	entry  map[string]chan bool
	entryV map[string]interface{}
	sync.RWMutex
}

func (sm *SyncMap) Out(key string, val interface{}) {

	sm.Lock()
	defer sm.Unlock()
	fmt.Printf("set value :%v\n", val.(string))
	sm.entryV[key] = val
	vchan := make(chan bool)
	sm.entry[key] = vchan
	//chan 里面有数据了
	go func() {
		vchan <- true
	}()
	fmt.Println("done!")

}
func (sm *SyncMap) Rd(key string, timeout time.Duration) interface{} {
	sm.Lock()
	defer sm.Unlock()
	//先判断存在不，不存在直接等待
	if v, ok := sm.entryV[key]; ok {
		return v
	} else {
		for {
			if sm.entry[key] == nil {
				continue
			}
			select {

			case _, ok := <-sm.entry[key]:
				//如果ok说明有chan,如果有chan那一定有值
				if ok {
					return sm.entryV[key]
				} else { //chan 关闭的情况直接跳出去，让他for 循环去
					break
				}
			case <-time.After(timeout):
				return nil
			}
		}
	}
}

func main() {
	syncmap := &SyncMap{
		entry:  make(map[string]chan bool),
		entryV: make(map[string]interface{}),
	}
	var s interface{} = "22"
	syncmap.Out("haha", s)
	syncmap.Out("haha", "33")
	syncmap.Out("haha", "44")

	println(syncmap.Rd("haha", 4*time.Second).(string))
	go func() { print(syncmap.Rd("zz", 3*time.Second).(string)) }()
	syncmap.Out("zz", "66")
}
