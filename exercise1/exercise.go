package exercise1

import "fmt"

func routine() {
	fmt.Printf("main function")
}

func SubMain() {
	go func() {
		fmt.Println("Hello World!")
	}()

	go routine()
}
