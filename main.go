package main

import (
	"fmt"
	"strings"
)

var LowercaseDictionary string = "abcdefghijklmnopqrstuvwxyz"
var UppercaseDictionary string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

const DictionaryLength = 26

type Encoder struct {
}

func NewEncoder(e *Encoder) *Encoder {
	return e
}

func (e *Encoder) Encrypt(str string, key int32) string {
	runes := []rune(str)
	new_str := []rune("")
	for i := 0; i < len(runes); i++ {
		IsLowercase := strings.ContainsRune(LowercaseDictionary, runes[i])
		IsUppercase := strings.ContainsRune(UppercaseDictionary, runes[i])
		if IsLowercase {
			if runes[i]+key > 'z' {
				new_str = append(new_str, runes[i]+key-DictionaryLength)
			} else {
				new_str = append(new_str, runes[i]+key)
			}
		} else {
			if IsUppercase {
				if runes[i]+key > 'Z' {
					new_str = append(new_str, runes[i]+key-DictionaryLength)
				} else {
					new_str = append(new_str, runes[i]+key)
				}
			} else {
				new_str = append(new_str, runes[i])
			}
		}
	}
	return string(new_str)
}

func (e *Encoder) Decrypt(str string, key int32) string {
	runes := []rune(str)
	new_str := []rune("")
	for i := 0; i < len(runes); i++ {
		IsLowercase := strings.ContainsRune(LowercaseDictionary, runes[i])
		IsUppercase := strings.ContainsRune(UppercaseDictionary, runes[i])
		if IsLowercase {
			if runes[i]+key < 'a' {
				new_str = append(new_str, runes[i]-key-DictionaryLength)
			} else {
				new_str = append(new_str, runes[i]-key)
			}
		} else {
			if IsUppercase {
				if runes[i]+key < 'A' {
					new_str = append(new_str, runes[i]-key-DictionaryLength)
				} else {
					new_str = append(new_str, runes[i]-key)
				}
			} else {
				new_str = append(new_str, runes[i])
			}
		}
	}
	return string(new_str)
}

func main() {
	var a Encoder
	a1 := NewEncoder(&a)
	b := a1.Encrypt("Hello, world!", 3)
	fmt.Println(b)
	fmt.Println(a1.Decrypt(b, 3))
}
