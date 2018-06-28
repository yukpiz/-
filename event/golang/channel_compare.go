package main

import "fmt"

func main() {
	compare()
}

func compare() {
	//
	// slice の 比較
	//
	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	//fmt.Println(s1 == s2)  // コンパイルエラー. invalid operation: s1 == s2 (slice can only be compared to nil)
	fmt.Println(s1 == nil) // false (nil との 比較は可能)
	compareInterface(s1, s2)

	//
	// map の 比較
	//
	m1 := map[string]int{"a": 1}
	m2 := map[string]int{"a": 1}
	//fmt.Println(m1 == m2)  // コンパイルエラー. invalid operation: m1 == m2 (map can only be compared to nil)
	fmt.Println(m1 == nil) // false (nil との 比較は可能)
	compareInterface(m1, m2)

	//
	// func の 比較
	//
	f1 := func() {}
	f2 := func() {}
	//fmt.Println(f1 == f2)  // コンパイルエラー. invalid operation: f1 == f2 (func can only be compared to nil)
	fmt.Println(f1 == nil) // false (nil との 比較は可能)
	compareInterface(f1, f2)

	//
	// chan の 比較. [Comparing Values in Go](https://medium.com/learning-the-go-programming-language/comparing-values-in-go-8f7b002e767a) 参照
	//
	ch0 := make(chan int)
	ch1 := make(chan int)
	ch2 := ch1
	fmt.Println(ch0 == ch1)    // false
	fmt.Println(ch1 == ch2)    // true
	fmt.Println(ch0 == nil)    // false (nil との 比較)
	compareInterface(ch0, ch1) // false
	compareInterface(ch1, ch2) // true

	m3 := make(map[chan int]int) // map の key として
	m3[ch1] = 1
	fmt.Println(m3[ch1]) // 1
	fmt.Println(m3[ch2]) // 1
	fmt.Println(m3[ch0]) // 0
}

func compareInterface(x, y interface{}) {
	defer func() {
		if p := recover(); p != nil {
			fmt.Printf("【%T 型 の 比較】で【panic 発生】\n", x)
		}
	}()

	fmt.Println(x == y)
}
