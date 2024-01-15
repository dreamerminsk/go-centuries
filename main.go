package main

import (
	"strconv"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetPlayer(text string) string {
	player := strings.TrimSpace(text)
	if strings.Contains(player, "[[") {
		player = strings.Replace(player, "[[", "", 1)
	}
	if strings.Contains(player, "]]") {
		player = strings.Replace(player, "]]", "", 1)
	}
	if strings.Contains(player, "|") {
		player = strings.Split(player, "|")[0]
	}
	return player
}
func GetCenturies(text string) []int{
	centuries:= []string{}
	values :=strings.Split(strings.TrimeSpace(text), ",")
	for _, value := range centuries {
	century, err := strconv.Atoi(strings.TrimSpace(value))
if err != nil {
	if strings.Contains(value, "147") {
		cenruries=append(centuries, 147)
		}
} else {
		cenruries=append(centuries, century)
}

}
	return centuries
}
func main() {

	file, _ := os.Open("./raw/2024 Masters (snooker).wiki")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "*") {
			if strings.Contains(line, "{{ndash}}") {
				parts := strings.Split(line, "{{ndash}}")
				fmt.Println(GetPlayer(parts[1]))
				fmt.Println(parts[0])
			}
		}
	}

}
