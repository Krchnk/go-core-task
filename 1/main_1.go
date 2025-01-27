package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func printType(v interface{}) string {
	return reflect.TypeOf(v).String()
}

func toString(v interface{}) string {
	switch v := v.(type) {
	case int:
		return strconv.Itoa(v)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case string:
		return v
	case bool:
		return strconv.FormatBool(v)
	case complex64:
		return fmt.Sprintf("%v", v)
	default:
		return ""
	}
}

func concatenateVariables(vars ...interface{}) string {
	var result strings.Builder
	for _, v := range vars {
		result.WriteString(toString(v))
		result.WriteString(" ")
	}
	return result.String()
}

func stringToRunes(s string) []rune {
	return []rune(s)
}

func hashRunes(runes []rune, salt string) string {
	str := string(runes)
	mid := len(str) / 2
	saltedStr := str[:mid] + salt + str[mid:]

	hash := sha256.Sum256([]byte(saltedStr))

	return hex.EncodeToString(hash[:])
}

func main() {

	var numDecimal int = 42
	var numOctal int = 052
	var numHexadecimal int = 0x2A
	var pi float64 = 3.14
	var name string = "Golang"
	var isActive bool = true
	var complexNum complex64 = 1 + 2i

	fmt.Println("Типы переменных:")
	fmt.Println("numDecimal:", printType(numDecimal))
	fmt.Println("numOctal:", printType(numOctal))
	fmt.Println("numHexadecimal:", printType(numHexadecimal))
	fmt.Println("pi:", printType(pi))
	fmt.Println("name:", printType(name))
	fmt.Println("isActive:", printType(isActive))
	fmt.Println("complexNum:", printType(complexNum))

	combinedString := concatenateVariables(numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum)
	fmt.Println("\nОбъединенная строка:", combinedString)

	runes := stringToRunes(combinedString)
	fmt.Println("\nСрез рун:", runes)

	salt := "go-2024"
	hashedResult := hashRunes(runes, salt)
	fmt.Println("\nХэшированный результат:", hashedResult)
}
