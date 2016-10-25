// channelを使ったgoroutineのサンプル
package main

import "fmt"

func sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	c <- total // totalをチャンネルへ送信する
}

func main() {
	a := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(a[:], c)
	go sum(a[:len(a)/2], c)

	x, y := <-c, <-c // チャンネルから受信する, ここでブロックする
	fmt.Println(x)
	fmt.Println(y)
}
