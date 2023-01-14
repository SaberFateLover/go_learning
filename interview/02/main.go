package main

import "fmt"

//给一个字符串判断每个字符都不相同

type dummy struct{}

func main() {
	//很简单的去重问题

	//只需要一个set即可

	fmt.Println(5 % 2)

}

//使用了额外的存储结构
func checkCharRepeat(s string) bool {
	set := make(map[byte]dummy)
	for v := range s {
		set[byte(v)] = dummy{}
	}
	if len(s) == len(set) {
		return true
	}
	return false
}
