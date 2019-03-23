package main

import "fmt"

func main()  {
	defer fmt.Println("Fourth")
	fmt.Println("First")
	fmt.Println("Third")
}