// 複数のチャンネルを扱う差異にselectでチャンネルを受信することでブロックすることが出来る
// Goの有名なイディオム
package main

import (
	"fmt"
	"time"
)

func fibo(c chan int, quit chan bool) {
	x, y := 1, 1
	for {
		select {
		// cチャンネルに送ってる間は処理を継続する
		case c <- x:
			x, y = y, x+y
		// quitチャンネルから何か受信したら処理終了
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan bool)
	// ブロックして結果を表示するgooutine
	go func() {
		for i := 0; i < 10; i++ {
			// fibo()の結果をブロックして表示
			fmt.Println(<-c)
			time.Sleep(time.Second)
		}
		// 指定回数分処理が終わればquitチャンネルに送信
		quit <- true
	}()
	fibo(c, quit) // 関数呼び出し
}
