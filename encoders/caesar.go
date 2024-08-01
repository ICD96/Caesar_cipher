package encoders

import (
	"errors"
)

type Caesar struct {
	lowerList string
	upperList string
}

func NewEncoder(lowerList, upperList string) (*Caesar, error) {
	if len(lowerList) != len(upperList) {
		return &Caesar{lowerList: lowerList, upperList: upperList}, errors.New("len of lowerList not equal to len of upperList")
	}
	return &Caesar{lowerList: lowerList, upperList: upperList}, nil
}

func (e *Caesar) ValidEncrypt(str string, shift int) (string, error) {
	err := e.validateShift(int32(shift))
	if err != nil {
		return "", err
	}
	return e.encrypt(str, int32(shift)), err
}

func (e *Caesar) validateShift(shift int32) error {
	if shift == 0 {
		return errors.New("shift cannot be 0")
	}
	if shift < 0 {
		return errors.New("shift cannot be a negative number")
	}
	if shift == int32(len(e.lowerList)) {
		return errors.New("shift is equal to alphabet len")
	}
	if shift > int32(len(e.lowerList)) {
		return errors.New("shift is higher then alphabet len")
	}
	return nil

}

func (e *Caesar) encrypt(str string, shift int32) string {
	runes := []rune(str)
	encryptedMessage := []rune{}
	for i := 0; i < len(runes); i++ {
		if runes[i]+shift > int32(e.lowerList[len(e.lowerList)-1]) {
			encryptedMessage = append(encryptedMessage, runes[i]+shift-int32(len(e.lowerList)))
		} else {
			encryptedMessage = append(encryptedMessage, runes[i]+shift)
		}
	}
	return string(encryptedMessage)
}

func (e *Caesar) ValidDecrypt(str string, shift int) (string, error) {
	err := e.validateShift(int32(shift))
	if err != nil {
		return "", err
	}
	return e.decrypt(str, int32(shift)), err
}

func (e *Caesar) decrypt(str string, shift int32) string {
	runes := []rune(str)
	decryptedMessage := []rune{}
	for i := 0; i < len(runes); i++ {
		if runes[i]-shift < int32(e.lowerList[0]) {
			decryptedMessage = append(decryptedMessage, runes[i]-shift+int32(len(e.lowerList)))
		} else {
			decryptedMessage = append(decryptedMessage, runes[i]-shift)
		}
	}
	return string(decryptedMessage)
}
