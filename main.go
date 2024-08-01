package main

import (
	"fmt"

	"example.com/encoders"
	"example.com/leteris"

	"github.com/spf13/pflag"
)

func main() {
	encoder, err := encoders.NewEncoder(leteris.LowerEnglishLet, leteris.UpperEnglishLet)
	if err != nil {
		fmt.Println("encoder err", err)
		return
	}

	var encryptedMessage, decryptedMessage string
	var key int

	pflag.StringVarP(&encryptedMessage, "encrypt", "e", "", "Шифрование соообщения")
	pflag.StringVarP(&decryptedMessage, "decrypt", "d", "", "Дешифрование соообщения")
	pflag.IntVarP(&key, "key", "k", 0, "Ключ сдвига при шифровании")

	pflag.Parse()

	//TODO here
	encoders.NewCaesarEncoder()
	if len(encryptedMessage) != 0 {
		fmt.Println()
	}
}
