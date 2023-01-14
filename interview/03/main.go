package main

import "fmt"

//翻转一个字符串
//字符串翻转有技巧可言哈
func main() {

	//学到了
	a := []int{1, 2, 4, 5, 6, 7}
	for i := 0; i < len(a)/2; i++ {
		a[i], a[len(a)-i-1] = a[len(a)-i-1], a[i]
	}
	fmt.Printf("%v", a)

	// fmt.Println(reverse("abcdefghijk"))
	// fmt.Println(reverse("abcde"))
	// fmt.Println(reverse("abc"))
	// fmt.Println(reverse("ab"))
	// fmt.Println(reverse("a"))
	//
	//
	//fmt.Println(rev("abcdefghijk"))
	//fmt.Println(rev("abcde"))
	//fmt.Println(rev("abc"))
	//fmt.Println(rev("ab"))
	//fmt.Println(rev("a"))
}

func reverse(s string) string {
	if len(s) == 1 {
		return string(s[0])
	}
	if len(s) == 2 {
		return string(s[1]) + string(s[0])
	}
	if len(s) == 3 {
		return string(s[2]) + string(s[1]) + string(s[0])
	}
	if len(s)%2 == 0 {
		return reverse(s)
	} else {
		return reverse(s[len(s)/2+1:]) + string(s[len(s)/2]) + reverse(s[:len(s)/2])
	}
}

//python ,go 都有的语法，学到了

// a,b =b,a 直接就能交换,字符串好像不行，但是数组可以😄
func rev(s string) string {
	//😄直接转数组
	chars := []rune(s)
	l := len(s)
	for i := 0; i < len(chars)/2; i++ {
		chars[i], chars[l-i-1] = chars[l-i-1], chars[i]
	}
	return string(chars)
}
