package calculator

import (
	"strconv"
)

func in(arr []rune, val rune) bool {
	for _, el := range arr {
		if el == val {
			return true
		}
	}
	return false
}

func Tokenize(expression string) ([]iToken, error) {
	l := len(expression)
	//nums := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	operations := []rune{'+', '-', '*', '/'}
	brackets := []rune{'(', ')'}

	result := make([]iToken, 0)

	var i, j int
	for i < l {
		if in(operations, rune(expression[i])) {
			result = append(result, opToken{value: rune(expression[i])})
		} else if expression[i] >= '0' && expression[i] <= '9' {
			j = i + 1
			if j < l {
				for j < l && expression[j] >= '0' && expression[j] <= '9' || expression[j] == '.' {
					j++
					if j == l {
						break
					}
				}
			}
			num, err := strconv.ParseFloat(string(expression[i:j]), 64)
			if err != nil {
				return make([]iToken, 0), err
			}
			i = j - 1
			result = append(result, numToken{value: num})
		} else if in(brackets, rune(expression[i])) {
			result = append(result, brToken{value: rune(expression[i])})
		}
		i++
	}
	return result, nil
}
