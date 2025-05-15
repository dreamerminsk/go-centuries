package main

import (
	"fmt"
	"github.com/trietmn/go-wiki"
)

func main() {
	gowiki.SetLanguage("en")

	page, err := gowiki.GetPage("2025 World Snooker Championship", -1, false, true)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		return
	}

	content, err := page.GetSection("Main draw")
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		return
	}

	fmt.Println(content)
}