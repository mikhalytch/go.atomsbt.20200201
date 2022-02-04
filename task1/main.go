package main

import "fmt"

func main() {
	fmt.Println(Shout(*Create(DogType), *Create(CatType)))
}