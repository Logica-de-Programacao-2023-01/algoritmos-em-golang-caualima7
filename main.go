package main

import "fmt"

func main() {
	var name string
	fmt.Print("qual seu nome ")
	fmt.Scan(&name)
	fmt.Println("olá,", name)
}
