// チャンネルをタイムアウトさせる方法
// 1.7.1?以降はcontextで行うのがナウいっぽい
package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	o := make(chan bool)

	go func() {
		for {
			select {
			// cから受信できればprint
			case v := <-c:
				fmt.Println(v)
			// 何も受信されない状態が5秒続くとこいつが発火
			case <-time.After(5 * time.Second):
				fmt.Println("timeout")
				o <- true
				break // ループ抜けるのでgoroutine終了
			}
		}
	}()
	c <- 1 // cに送る
	c <- 2 // cに送る
	<-o    // ここでブロック
}
