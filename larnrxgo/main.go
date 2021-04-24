// 一个异步编程的 API

package main

import (
	"fmt"

	"github.com/reactivex/rxgo"
)

func main() {
	observable := rxgo.Just(1, 2, 3, 4, 5)()
	ch := observable.Observe()
	for item := range ch {
		fmt.Println(item.V)
	}
}
