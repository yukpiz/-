package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x: //ここでc送信
			x, y = y, x+y
		case <-quit: //quit受信
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c) //ここでc受信
		}
		quit <- 0 //ここでquit送信
		//quit <- struct{}() //よくある シグナル用ですよ
	}()
	fibonacci(c, quit)
}
