package main

import (
	"fmt"
	"regexp"
)

var Dictionary string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type Encoder struct {
}

func NewEncoder() {
}

func (e *Encoder) Encrypt(str string, key int32) string {
	runes := []rune(str)
	new_str := []rune("")
	for i := 0; i < len(runes); i++ {
		Isletter := regexp.MustCompile(`^[a-zA-Z]+$`).MatchString(string(runes[i]))
		Iscapital := regexp.MustCompile(`^[A-Z]+$`).MatchString(string(runes[i]))
		if Isletter {
			if Iscapital {
				if runes[i]+key > 90 {
					new_str = append(new_str, runes[i]+key-26)
				} else {
					new_str = append(new_str, runes[i]+key)
				}
			} else {
				if runes[i]+3 > 122 {
					new_str = append(new_str, runes[i]+key-26)
				} else {
					new_str = append(new_str, runes[i]+key)
				}
			}
		} else {
			new_str = append(new_str, runes[i])
		}
	}
	return string(new_str)
}

func (e *Encoder) Decrypt(str string, key int32) string {
	runes := []rune(str)
	new_str := []rune("")
	for i := 0; i < len(runes); i++ {
		Isletter := regexp.MustCompile(`^[a-zA-Z]+$`).MatchString(string(runes[i]))
		Iscapital := regexp.MustCompile(`^[A-Z]+$`).MatchString(string(runes[i]))
		if Isletter {
			if Iscapital {
				if runes[i]-key < 65 {
					new_str = append(new_str, runes[i]-key+26)
				} else {
					new_str = append(new_str, runes[i]-key)
				}
			} else {
				if runes[i]-key < 97 {
					new_str = append(new_str, runes[i]-key+26)
				} else {
					new_str = append(new_str, runes[i]-key)
				}
			}
		} else {
			new_str = append(new_str, runes[i])
		}
	}
	return string(new_str)
}

func main() {
	var a Encoder
	b := a.Encrypt("Hello, world!", 3)
	fmt.Println(b)
	fmt.Println(a.Decrypt(b, 3))
}
