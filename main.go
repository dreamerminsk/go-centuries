package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type PlayerStats struct {
	Player    string
	Centuries []int
}

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

func GetCenturies(text string) []int {
	centuries := []int{}
	values := strings.Split(strings.TrimSpace(strings.Replace(text, "*", "", 1)), ",")
	for _, value := range values {
		century, err := strconv.Atoi(strings.TrimSpace(value))
		if err != nil {
			if strings.Contains(value, "147") {
				centuries = append(centuries, 147)
			}
		} else {
			centuries = append(centuries, century)
		}

	}
	return centuries
}

func GetPlayerStats(value string) PlayerStats {
	if strings.Contains(line, "{{ndash}}") {
				parts := strings.Split(line, "{{ndash}}")
				centuries := GetCenturies(parts[0])
				player := GetPlayer(parts[1])
		return  PlayerStats{player,  centuries}
			}
return nil

}

func ProcessFile(name string) {
	file, _ := os.Open(name)
	scanner := bufio.NewScanner(file)
	ps := 0
	cs := 0
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "*") {
			if strings.Contains(line, "{{ndash}}") {
				parts := strings.Split(line, "{{ndash}}")
				centuries := GetCenturies(parts[0])
				player := GetPlayer(parts[1])
				ps = ps + 1
				cs = cs + len(centuries)
				fmt.Println(player, len(centuries))
				for _, century := range centuries {
					fmt.Println("\t", century)
				}
			}
		}
	}
	fmt.Println("players: ", ps, ", ", "centuries: ", cs)
}

func main() {
	files, err := ioutil.ReadDir(filepath.Join(".", "raw"))
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if !file.IsDir() {
			fmt.Println(file.Name())
			ProcessFile(filepath.Join(".", "raw", file.Name()))
		}
	}
}
