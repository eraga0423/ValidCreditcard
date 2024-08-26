package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Получает флаг и отделяет что написано после "=" и  ее отправляет

func Readflag(flag string) string {
	res := ""
	for ind, val := range flag {
		if val == '=' {
			res = flag[ind+1:]
		}
	}
	return res
}

// читает файл, разбивает строки на ключ и значение, сохраняет их в карту (map), и возвращает её.
// В случае ошибки чтения файла программа завершает работу.

func Readfiles(file string) map[string]string {
	var brand string
	var num string
	data, err := os.Open(file)
	if err != nil {
		fmt.Println("Ошибка чтение файла", err)
		os.Exit(1)
	}
	defer data.Close()
	mapa := make(map[string]string)
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		line := scanner.Text()
		br := strings.Split(line, ":")
		if len(br) == 2 {
			brand = br[0]
			num = br[1]
		}
		mapa[num] = brand

	}
	return mapa
}
