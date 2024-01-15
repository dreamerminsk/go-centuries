package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	file, _ := os.Open("./raw/2024 Masters (snooker).wiki")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}
