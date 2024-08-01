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

	var messageToEncrypt, messageToDecrypt, encryptedMessage, decryptedMessage string
	var shift int

	pflag.StringVarP(&messageToEncrypt, "encrypt", "e", "", "Шифрование соообщения по ключу")
	pflag.StringVarP(&messageToDecrypt, "decrypt", "d", "", "Дешифрование соообщения по ключу")
	pflag.IntVarP(&shift, "shift", "s", 0, "Ключ сдвига при шифровании")
	pflag.Parse()

	if len(messageToEncrypt) != 0 {
		encryptedMessage, err = encoder.ValidEncrypt(messageToEncrypt, shift)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(encryptedMessage)
	}
	if len(messageToDecrypt) != 0 {
		decryptedMessage, err = encoder.ValidDecrypt(messageToDecrypt, shift)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(decryptedMessage)
	}
}
