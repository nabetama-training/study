package main

import "fmt"

func main() {
	ch := make(chan int, 2) // 2つのintの要素を保持できるチャンネル
	ch <- 2
	ch <- 3
	// ch <- 4 // error
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
