package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/pflag"
)

type Encoder struct {
	lowerList string
	upperList string
}

func NewEncoder(lowerList, upperList string) (*Encoder, error) {
	if len(lowerList) != len(upperList) {
		return &Encoder{lowerList: lowerList, upperList: upperList}, errors.New("len of lowerList not equal to len of upperList")
	}
	return &Encoder{lowerList: lowerList, upperList: upperList}, nil
}

func (e *Encoder) Encrypt(str string, key int32) string {
	runes := []rune(str)
	newStr := []rune{}
	for i := 0; i < len(runes); i++ {
		isLower := strings.ContainsRune(e.lowerList, runes[i])
		isUpper := strings.ContainsRune(e.upperList, runes[i])
		if isLower {
			if runes[i]+key > 'z' {
				newStr = append(newStr, runes[i]+key-int32(len(e.lowerList)))
			} else {
				newStr = append(newStr, runes[i]+key)
			}
		}
		if isUpper {
			if runes[i]+key > 'Z' {
				newStr = append(newStr, runes[i]+key-int32(len(e.upperList)))
			} else {
				newStr = append(newStr, runes[i]+key)
			}
		}
	}
	return string(newStr)
}

func (e *Encoder) Decrypt(str string, key int32) string {
	runes := []rune(str)
	newStr := []rune{}
	for i := 0; i < len(runes); i++ {
		isLower := strings.ContainsRune(e.lowerList, runes[i])
		isUpper := strings.ContainsRune(e.upperList, runes[i])
		if isLower {
			if runes[i]+key > 'a' {
				newStr = append(newStr, runes[i]-key+int32(len(e.lowerList)))
			} else {
				newStr = append(newStr, runes[i]-key)
			}
		}
		if isUpper {
			if runes[i]+key > 'A' {
				newStr = append(newStr, runes[i]-key+int32(len(e.upperList)))
			} else {
				newStr = append(newStr, runes[i]-key)
			}
		}
	}
	return string(newStr)
}

func main() {
	encoder, err := NewEncoder("abcdefghijklmnopqrstuvwxyz", "ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
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
