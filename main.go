package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"slices"
	"strings"
)

// Эта функция обрабатывает команды командной строки для валидации, генерации, получения информации и выпуска номеров карт.
// Она проверяет корректность ввода, выполняет нужные действия на основе команд, и завершает работу с ошибкой, если ввод некорректен.
func main() {
	var cardNum []string
	comand := []string{"validate", "generate", "information", "issue"}
	if os.Args == nil || len(os.Args) == 1 || !slices.Contains(comand, os.Args[1]) {

		fmt.Println("Ошибка ввода в терминал")
		os.Exit(1)
	}
	scann := bufio.NewScanner(os.Stdin)

	if os.Args[1] == "validate" {
		if slices.Contains(os.Args, "--stdin") && os.Args[1] == "validate" && len(os.Args) > 2 {
			for scann.Scan() {
				numCard := []string{}
				numCard = append(numCard, strings.Fields(scann.Text())...)
				Validate(numCard)
			}
		} else {
			cardNum = os.Args[2:]
			Validate(cardNum)
		}
	} else if err := scann.Err(); err != nil {
		fmt.Println("ОШИБКА")
		os.Exit(1)
	}
	if os.Args[1] == "generate" && len(os.Args) > 1 {
		if slices.Contains(os.Args, "--pick") && os.Args[1] == "generate" && len(os.Args) > 2 {
			cardNum = os.Args[3:]
			arr := Generate(cardNum)
			ind := rand.Intn(len(arr))
			fmt.Println(arr[ind])
		} else {
			cardNum = os.Args[2:]
			for _, val := range Generate(cardNum) {
				fmt.Println(val)
			}
		}
	}
	count := 0
	brands := ""
	issuers := ""
	issue := ""
	brand := ""
	var filebrand map[string]string
	var fileissue map[string]string
	for _, flags := range os.Args {
		if strings.HasPrefix(flags, "--brands=") {
			brands = Readflag(flags)
			filebrand = Readfiles(Readflag(flags))
			count++

		}
		if strings.HasPrefix(flags, "--issuers=") {
			issuers = Readflag(flags)
			fileissue = Readfiles(Readflag(flags))
			count++

		}
		if strings.HasPrefix(flags, "--issuer=") {
			issue = Readflag(flags)
			count++
		}
		if strings.HasPrefix(flags, "--brand=") {
			brand = Readflag(flags)
			count++

		}

	}
	if os.Args[1] == "information" && count == 2 {
		if slices.Contains(os.Args, "--stdin") && os.Args[1] == "information" && len(os.Args) > 4 && count == 2 {
			for scann.Scan() {
				numCard := []string{}
				cardNum = append(numCard, strings.Fields(scann.Text())...)
				if len(cardNum) == 0 {
					fmt.Println("Ошибка ввода")
					os.Exit(1)
				}

				Information(brands, issuers, cardNum)

			}
		} else {
			cardNum := os.Args[4:]

			Information(brands, issuers, cardNum)
		}
	} else if (len(os.Args) < 5 || count < 2) && os.Args[1] == "information" {
		fmt.Println("Флаги указаны неправильно")
		os.Exit(1)
	}

	if os.Args[1] == "issue" && len(os.Args) == 6 && count == 4 {
		Issue(brand, issue, fileissue, filebrand)
	} else if (len(os.Args) < 6 || count < 4 || len(os.Args) > 6) && os.Args[1] == "issue" {
		fmt.Println("Флаги указаны неправильно")
		os.Exit(1)
	}
}
