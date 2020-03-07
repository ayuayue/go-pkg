package main

import (
	"fmt"
	"get_input/get_input"
)

func main() {
	str := get_input.GetString("name: ")
	fmt.Println(str)
	inter := get_input.GetInt("age: ")

	fmt.Println(inter)
}
