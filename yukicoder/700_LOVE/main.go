package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	row, _ := strconv.Atoi(strings.Split(nextLine(), " ")[0])

	for i := 0; i < row; i++ {
		s := nextLine()
		if strings.Index(s, "LOVE") >= 0 {
			success()
			os.Exit(0)
		}
	}
	fail()
}

func nextLine() string {
	if sc.Scan() {
		return sc.Text()
	}
	return ""
}

func success() {
	fmt.Println("YES")
}

func fail() {
	fmt.Println("NO")
}
