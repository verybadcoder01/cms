package internal

import (
	"strconv"
	"strings"
)

// DbNotationToArray принимает строку вида "число1,число2,..." и возвращает массив вида [число1, число2, ...]
func DbNotationToArray(DbNotation string) []int {
	var result []int
	numbers := strings.Split(DbNotation, ",")
	for _, number := range numbers {
		num, _ := strconv.Atoi(number)
		result = append(result, num)
	}
	return result
}

// IntArrayToDbNotation функция, обратная к предыдущей
func IntArrayToDbNotation(Data []int) string {
	if len(Data) == 0 {
		return ""
	}
	var result string
	for _, number := range Data {
		result += strconv.Itoa(number)
		result += ","
	}
	result = result[:len(result)-1]
	return result
}