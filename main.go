package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	file, _ := os.Open("./raw/2024 Masters (snooker).wiki")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "*") {
			if strings.Contains(line, "{{ndash}}") {
			parts := strings.Split(line, "{{ndash}}")
			fmt.Println(parts[1])
			fmt.Println(parts[0])
				}
		}
	}

}
