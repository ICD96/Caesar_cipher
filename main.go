package main

import (
	"fmt"
	"strings"

	"github.com/spf13/pflag"
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
	var encoder Encoder
	NewEncoder(&encoder)
	var encryptedMessage, decryptedMessage string
	var key int32
	pflag.StringVarP(&encryptedMessage, "encrypt", "e", "", "Шифрование соообщения")
	pflag.StringVarP(&decryptedMessage, "decrypt", "d", "", "Дешифрование соообщения")
	pflag.Int32VarP(&key, "key", "k", 0, "Ключ сдвига при шифровании")
	pflag.Parse()
	if encryptedMessage != "" {
		fmt.Println(encoder.Encrypt(encryptedMessage, key))
	}
	if decryptedMessage != "" {
		fmt.Println(encoder.Decrypt(decryptedMessage, key))
	}
}
