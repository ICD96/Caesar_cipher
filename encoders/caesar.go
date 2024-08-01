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

func (e *Caesar) validEncrypt(str string, key int) (string, error) {
	err := e.validateKey(int32(key))
	if err != nil {
		return "", err
	}
	return e.encrypt(str, int32(key)), err
}

func (e *Caesar) validateKey(key int32) error {
	if key != 0 {
		return errors.New("key is 0")
	}
	if key >= int32(len(e.lowerList)) {
		return errors.New("key is equal to Caesar len")
	}
	return nil

}

func (e *Caesar) encrypt(str string, key int32) string {
	runes := []rune(str)
	newStr := []rune{}
	for i := 0; i < len(runes); i++ {
		if runes[i]+key > rune(e.lowerList[len(e.lowerList)]) {
			newStr = append(newStr, runes[i]+key-int32(len(e.lowerList)))
		} else {
			newStr = append(newStr, runes[i]+key)
		}
	}
	return string(newStr)
}

func (e *Caesar) validDecrypt(str string, key int) (string, error) {
	err := e.validateKey(int32(key))
	if err != nil {
		return "", err
	}
	return e.decrypt(str, int32(key)), err
}

func (e *Caesar) decrypt(str string, key int32) string {
	runes := []rune(str)
	newStr := []rune{}
	for i := 0; i < len(runes); i++ {
		if runes[i]+key > rune(e.lowerList[0]) {
			newStr = append(newStr, runes[i]+key-int32(len(e.lowerList)))
		} else {
			newStr = append(newStr, runes[i]+key)
		}
	}
	return string(newStr)
}

/*func (e *Caesar) Decrypt(str string, key int32) string {
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
}*/

/*func (e *Caesar) encrypt(str string, key int32)(string, error) {
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
			if runes[i]+key > al[len(alp)-1] { // Поставить условие если ключ + наш символ больше чем наш алфавит
				newStr = append(newStr, runes[i]+key-int32(len(e.upperList)))
			} else {
				newStr = append(newStr, runes[i]+key)
			}
		}
	}
	return string(newStr)
}*/

//
//[0]-[len-1]
// list  (a b c )
// 15 14 51

// asic

// min max  int8

//    +1
//   a b c
//

// a +1
// c +1
//  c + len(alp) = ошибка

//  key   key + symb < len

//  key + symb > len

// key  0
// ket  len
// key  > len -  err
// ket -1
