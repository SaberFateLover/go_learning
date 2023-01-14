package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strings"
	"syscall"
)

var banner = `
        _                  __                  __         __ 
  ___  (_)  __ _    ___   / / ___       ____  / /  ___ _ / /_
 (_-< / /  /  ' \  / _ \ / / / -_)     / __/ / _ \/ _  // __/
/___//_/  /_/_/_/ / .__//_/  \__/      \__/ /_//_/\_,_/ \__/ 
/_/
`

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

//简单写一下，后面再增加
type Content struct {
	Name string
	Data string
}
type username struct {
	n string
}

var name *username

func main() {
	print(banner)

	//signs := make(chan os.Signal,1)
	//signal.Notify(signs,syscall.SIGINT,syscall.SIGTERM)

	//stop := make(chan bool,1)

	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("err : ", err)
		return
	}
	defer conn.Close() // 关闭TCP连接

	inputReader := bufio.NewReader(os.Stdin)
	go receiveData(conn)
	//processSignal(signs, stop)
	for {
		//
		input, _ := inputReader.ReadString('\n') // 读取用户输入

		inputInfo := strings.Trim(input, "\r\n")

		if strings.ToUpper(inputInfo) == "Q" { // 如果输入q就退出
			return
		}
		req := Content{
			Name: randomGetUsername(5).n,
			Data: inputInfo,
		}
		marshal, err := json.Marshal(&req)
		if err != nil {
			fmt.Printf("%v 序列化有误\n", req)
			//重新发数据吧
			continue
		}
		_, err = conn.Write(marshal) // 发送数据
		if err != nil {
			return
		}

	}
}

func receiveData(conn net.Conn) {

	for {
		//这个是数组，使用的时候需要转成切片
		buf := [512]byte{}

		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("recv failed, err:", err)
			return
		}
		fmt.Println(string(buf[:n]))
		resp := &Content{}
		err = json.Unmarshal(buf[:n], resp)
		if err != nil {
			fmt.Println("unmarshal err:", err)
			return
		}
		fmt.Printf("response data name: %v, data : %v\n", resp.Name, resp.Data)
		return
	}
}

func processSignal(signs chan os.Signal, stop chan bool) {
	go func() {
		s := <-signs
		fmt.Println("signal: ", s)
		if s == syscall.SIGINT || s == syscall.SIGTERM {
			stop <- true
		}
	}()

	go func() {
		stop := <-stop
		if stop {
			fmt.Println("断开链接！！")
			os.Exit(1)
		}
	}()
}

//单例模式
func randomGetUsername(n int) *username {
	if name == nil {
		b := make([]byte, n)
		for i := range b {
			b[i] = letterBytes[rand.Intn(len(letterBytes))]
		}
		name = &username{
			n: string(b),
		}
		return name
	}
	return name
}
