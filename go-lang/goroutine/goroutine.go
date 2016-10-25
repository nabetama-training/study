// チャンネルを使わないgoroutine
package main

import (
	"fmt"
	"runtime"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched() // CPUの時間を受け渡す
		fmt.Println(s)
	}
}

func main() {
	go say("hello")
	say("world")
}
