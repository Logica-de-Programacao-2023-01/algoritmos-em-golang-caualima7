package main

import "fmt"

func main() {
	for i := 1; i < 11; i++ {
		if i == 11 {
			break
		}
		fmt.Println(i)
	}
	fmt.Println("fim do loop com break")
}
