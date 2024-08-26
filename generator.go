package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

// Эта функция генерирует уникальный номер кредитной карты на основе заданного бренда и выпуска, проверяя их соответствие по заданным картам (issuersmap и brandmap).
// Если соответствие не найдено, программа завершает работу с сообщением об ошибке.
// Если соответствие найдено, функция генерирует номер карты с использованием алгоритма Луна (Luhn's Algorithm) и выводит его на экран.

func Issue(inbrand string, inissue string, issuersmap map[string]string, brandmap map[string]string) {
	valIssue := ""
	valBrand := ""
	flag := false
	res := ""
	for num2, issuer2 := range issuersmap {
		if inissue == issuer2 {
			valIssue = num2
		}
	}
	for num1, brand := range brandmap {
		if strings.HasPrefix(valIssue, num1) && brand == inbrand {
			flag = true
			valBrand = brand
		}
	}
	if !flag {
		fmt.Println("Бренд не соответствует кредитной карточке банка")
		os.Exit(1)
	}
	for {
		if valBrand == "AMEX" {
			n := 100000000 + rand.Intn(999999999-100000000+1)
			val := valIssue + strconv.Itoa(n)
			if LuhnaAlgoritm(val) {
				res = val
				break
			}
		} else {
			n := 1000000000 + rand.Intn(9999999999-1000000000+1)
			val := valIssue + strconv.Itoa(n)
			if LuhnaAlgoritm(val) {
				res = val
				break
			}
		}
	}
	fmt.Println(res)
}

// Эта функция генерирует возможные варианты номеров карт, заменяя символы * на цифры и проверяя их по алгоритму Луна.
// Если количество или расположение * некорректно, функция завершает работу с ошибкой.

func Generate(num []string) []string {
	if len(num) == 0 {
		fmt.Println("Пустая строка")
		os.Exit(1)
	}
	count := 0
	numbe := strings.Join(num, ",")

	res := []string{}
	for _, number := range numbe {
		if number == '*' {
			count++
		}
	}
	if count > 4 || count == 0 {
		fmt.Println("Много звездочек либо их нет")
		os.Exit(1)
	}
	index := len(numbe) - count
	if count == 1 {
		if numbe[index] != '*' {
			fmt.Println("Звездочки не в том месте 1")
			os.Exit(1)
		}
		for i := 0; i < 10; i++ {
			a := strings.Replace(numbe, "*", strconv.Itoa(i), -1)
			if LuhnaAlgoritm(a) {
				res = append(res, a)
			}
		}
	}
	if count == 2 {
		if string(numbe[index:]) != "**" {
			fmt.Println("Звездочки не в том месте 2")
			os.Exit(1)
		}
		for i := 0; i < 10; i++ {
			a := strings.Replace(numbe, "*", strconv.Itoa(i), 1)
			for u := 0; u < 10; u++ {
				b := strings.Replace(a, "*", strconv.Itoa(u), 1)
				if LuhnaAlgoritm(b) {
					res = append(res, b)
				}
			}
		}
	}
	if count == 3 {
		if string(numbe[index:]) != "***" {
			fmt.Println("Звездочки не в том месте 3")
			os.Exit(1)
		}
		for i := 0; i < 10; i++ {
			a := strings.Replace(numbe, "*", strconv.Itoa(i), 1)
			for u := 0; u < 10; u++ {
				b := strings.Replace(a, "*", strconv.Itoa(u), 1)
				for o := 0; o < 10; o++ {
					c := strings.Replace(b, "*", strconv.Itoa(o), 1)
					if LuhnaAlgoritm(c) {
						res = append(res, c)
					}
				}
			}
		}
	}
	if count == 4 {
		if string(numbe[index:]) != "****" {
			fmt.Println("Звездочки не в том месте 4")
			os.Exit(1)
		}
		for i := 0; i < 10; i++ {
			a := strings.Replace(numbe, "*", strconv.Itoa(i), 1)
			for u := 0; u < 10; u++ {
				b := strings.Replace(a, "*", strconv.Itoa(u), 1)
				for o := 0; o < 10; o++ {
					c := strings.Replace(b, "*", strconv.Itoa(o), 1)
					for p := 0; p < 10; p++ {
						d := strings.Replace(c, "*", strconv.Itoa(p), 1)
						if LuhnaAlgoritm(d) {
							res = append(res, d)
						}
					}
				}
			}
		}
	}
	if len(res) == 0 {
		fmt.Println("Неправильный ввод")
		os.Exit(1)
	}
	return res
}
