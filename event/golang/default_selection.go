package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(100 * time.Millisecond)  //定期実行
	boom := time.After(500 * time.Millisecond) //s秒後に1度実行
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(20000 * time.Millisecond)
		}
	}
}
