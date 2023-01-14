package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"sync"
)

var banner = `
        _                  __                  __         __ 
  ___  (_)  __ _    ___   / / ___       ____  / /  ___ _ / /_
 (_-< / /  /  ' \  / _ \ / / / -_)     / __/ / _ \/ _  // __/
/___//_/  /_/_/_/ / .__//_/  \__/      \__/ /_//_/\_,_/ \__/ 
/_/
`

// Content 简单写一下，后面再增加
type Content struct {
	Name string
	Data string
}

type Conn struct {
	conns []net.Conn
	lock  sync.Mutex
}

var conns []net.Conn

func main() {

	print(banner)

	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Listen failed ", err)
		return
	}
	var c = make(chan net.Conn, 2)
	//一个协程监听
	go acceptConnect(listen, c)
	//再开个协程处理
	go process(c)

	//让server挂起来
	select {}
	//for{
	//	accept, err := listen.Accept()
	//	if err!=nil{
	//		fmt.Println("accept failed",err)
	//		return
	//	}
	//	conns=append(conns, accept)
	//	var c= make(chan net.Conn,1)
	//	c <-accept
	//	go process(c)
	//	//process(accept)
	//}

}

func acceptConnect(listen net.Listener, c chan net.Conn) {
	for {
		accept, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed", err)
			return
		}
		//
		var lock sync.Mutex
		lock.Lock()
		conns = append(conns, accept)
		lock.Unlock()
		//var c= make(chan net.Conn,1)
		c <- accept
		//go process(c)
		//process(accept)
	}

}

// TCP Server端测试
// 处理函数
func process(conns chan net.Conn) {
	conn := <-conns
	reader := bufio.NewReader(conn)
	//链接不要关闭
	for {
		var buf [128]byte
		n, err := reader.Read(buf[:]) // 读取数据
		resp := &Content{}
		if err != nil {
			fmt.Println("read from client failed, err: ", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("收到Client端发来的数据：", recvStr)
		err = json.Unmarshal(buf[:n], resp)
		if err != nil {
			fmt.Println("unmarshal err: ", err)
			break
		}
		fmt.Printf("response data name: %v, data : %v\n", resp.Name, resp.Data)

		//conn.Write([]byte(recvStr)) // 发送数据
		broadcastData([]byte(recvStr), conn)
	}
}

func broadcastData(b []byte, conn net.Conn) {
	for i := 0; i < len(conns); i++ {
		//if conns[i]== conn{
		//	continue
		//}
		_, _ = conns[i].Write(b)
	}

}
