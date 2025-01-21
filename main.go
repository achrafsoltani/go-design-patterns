package main

import "fmt"

func main() {
	fmt.Println("Singleton")
	for i := 0; i < 30; i++ {
		go GetInstance()
	}
}
