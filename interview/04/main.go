package main

import (
	"fmt"
	"sort"
	"strings"
)

type Chars []rune

func (s Chars) Len() int {
	return len(s)
}

func (s Chars) Less(i, j int) bool {
	return s[i] > s[j]
}

func (s Chars) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {

	print(checkSameStr("abc", "cba"))

}

func checkSameStr(s string, s2 string) bool {

	if len(s) != len(s2) {
		return false
	}

	char := Chars(s)
	sort.Sort(char)
	fmt.Printf("%v\n", char)
	char2 := Chars(s2)
	fmt.Printf("%v\n", char2)
	sort.Sort(char2)
	return string(char) == string(char2)

}

//别人的方法，试试看
func isRegroup(s1, s2 string) bool {
	sl1 := len([]rune(s1))
	sl2 := len([]rune(s2))

	if sl1 > 5000 || sl2 > 5000 || sl1 != sl2 {
		return false
	}

	for _, v := range s1 {
		if strings.Count(s1, string(v)) != strings.Count(s2, string(v)) {
			return false
		}
	}
	return true
}
