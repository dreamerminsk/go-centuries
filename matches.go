package main

import (
	"fmt"
	"github.com/trietmn/go-wiki"
)

func main() {
	wiki.SetLang("en")

	page, err := wiki.GetPage("2025 World Snooker Championship", -1, false, true)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		return
	}

	content, err := page.GetContent()
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		return
	}

	fmt.Println(content)
}