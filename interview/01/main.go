package main

import (
	"fmt"
	"sync"
)

//交替打印26个英文字母
//核心两个chan,加一个waitGroup
func main() {

	letter, letter2 := make(chan byte), make(chan byte)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func(chan byte, chan byte) {
		letter2 <- 'a'
		for {
			select {
			case word := <-letter:
				fmt.Printf("%c ", word)
				letter2 <- word + 1
				if word == 'z' {
					wg.Done()
				}
			}
		}

	}(letter, letter2)

	go func(chan byte, chan byte) {
		for {
			select {
			case word := <-letter2:
				fmt.Printf("%c ", word)
				if word == 'z' {
					wg.Done()
				}
				letter <- word + 1
			}
		}

	}(letter2, letter)

	wg.Wait()

}
