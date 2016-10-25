package main

import "fmt"

func fibo(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x // send!
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibo(cap(c), c)

	// range <chan> でこのチャンネルがclose(c)されるまでチャンネルのデータを読み込むことが出来る
	for i := range c {
		fmt.Println(i)
	}
}
