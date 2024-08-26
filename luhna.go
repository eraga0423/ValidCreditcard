package main

import (
	"strconv"
	"strings"
)

// Это функция проверяет цифры на Луна алгоритм а также на бренд карты

func LuhnaAlgoritm(num string) bool {
	sum := 0
	parity := len(num) % 2
	flag := false
	for o, chara := range num {
		brands := Readfiles("brands.txt")
		for num1 := range brands {
			if strings.HasPrefix(string(num), num1) && len(num) > 13 && len(num) < 19 {
				flag = true
			}
		}
		digi, error := strconv.Atoi(string(chara))
		if error != nil {
			return false
		}
		if o%2 == parity {
			digi *= 2
			if digi > 9 {
				digi -= 9
			}

		}
		sum += digi
	}
	return sum%10 == 0 && flag
}
