package exercise2

import (
	"fmt"
	"time"
)

func printO() {
	go func() {
		println(1)
	}()
}

func SubMain() {
	printO()
	time.Sleep(2 * time.Second)
	fmt.Println(0)
}
