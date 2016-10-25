// contextを用いたtimeoutのイディオム
// 1.7で実装された
package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// FooFoo 1/2でエラーを返す無意味な関数
func FooFoo() (string, error) {
	rand.Seed(time.Now().UnixNano())
	var bools = [2]error{errors.New("A"), nil}
	return "", bools[rand.Intn(2)]
}

// Foo なんかしてなんかする
func Foo(ctx context.Context) error {
	errCh := make(chan error, 1)
	go func() {
		_, err := FooFoo()
		if err != nil {
			errCh <- err
		}
	}()

	select {
	case err := <-errCh:
		if err != nil {
			return err
		}
	// DoneはContextを実行中の関数へのキャンセレーションシグナル(godoc)
	case <-ctx.Done():
		return ctx.Err()
	}
	return nil
}

func main() {
	// 2秒でタイムアウトするcontextを生成
	// cancelを実行することでTimeout前にキャンセルを実行することができる
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	errCh := make(chan error, 1)
	go func() {
		errCh <- Foo(ctx)
	}()

	select {
	case err := <-errCh:
		if err != nil {
			fmt.Println("failed: ", err)
			return
		}
	}
	fmt.Println("Normal Exit")
}
