package main

import (
	"fmt"
	"os"
	"strings"
)

// Проверяет из файла "brand.txt" и "issue.txt" наличие флагов если они существует отправляет на функцию "printinf"

func Information(brands string, issuers string, num []string) {
	issue := ""
	brand := ""
	if len(num) == 0 {
		fmt.Println("Не набрали номер карты")
		os.Exit(1)
	}
	for _, val := range num {

		for key1, value1 := range Readfiles(issuers) {
			if strings.HasPrefix(val, key1) {
				issue = value1
			}
		}
		for key2, value2 := range Readfiles(brands) {
			if strings.HasPrefix(val, key2) {
				brand = value2
			}
		}

		Printinf(issue, brand, val)
	}
}

//  Делает принт с "information" проверяя ее через Луна алгоритм

func Printinf(issue string, brand string, num string) {
	if !LuhnaAlgoritm(num) || issue == "" || brand == "" || num == "" {
		fmt.Printf("\n%s\nCorrect: no\nCard Brand: -\nCard Issuer: -\n\n", num)
		os.Exit(1)
	} else if LuhnaAlgoritm(num) {
		fmt.Printf("\n%s\nCorrect: yes\nCard Brand: %s\nCard Issuer: %s\n\n", num, brand, issue)
	}
}

// Проверяет валидность проверяем через Луна алгоритм

func Validate(num []string) {
	if len(num) == 0 {
		fmt.Println("INCORRECT")
		os.Exit(1)
	}
	for _, number := range num {
		if !LuhnaAlgoritm(number) {
			fmt.Println("INCORRECT")
			os.Exit(1)
		} else if LuhnaAlgoritm(number) {
			fmt.Println("OK")
		}
	}
}
