package main

import "fmt"

func main1() {
	runes := 'a'
	new_str := []rune("")
	new_str = append(new_str, runes+3)
	fmt.Println(string(new_str))
}
