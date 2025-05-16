package main

import (
	"bufio"
	"fmt"
	"github.com/dreamerminsk/go-centuries/wiki"
	"github.com/dreamerminsk/go-centuries/wiki/client"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type EventStats struct {
	Event     string
	Players   int
	Centuries int
}

type PlayerStats struct {
	Player    string
	Centuries []int
}

type Centuries struct {
	C100 int
	C110 int
	C120 int
	C130 int
	C140 int
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

func GetPlayerStats(value string) *PlayerStats {
	if strings.Contains(value, "{{ndash}}") {
		parts := strings.Split(value, "{{ndash}}")
		centuries := GetCenturies(parts[0])
		player := GetPlayer(parts[1])
		return &PlayerStats{Player: player, Centuries: centuries}
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
			pss := GetPlayerStats(line)
			if pss != nil {
				ps = ps + 1
				cs = cs + len(pss.Centuries)
				fmt.Println(pss.Player, len(pss.Centuries))
				for _, century := range pss.Centuries {
					fmt.Println("\t", century)
				}
			}
		}
	}
	fmt.Println("players: ", ps, ", ", "centuries: ", cs)
}

func ProcessFiles() {
	files, err := ioutil.ReadDir(filepath.Join(".", "2023–24", "raw"))
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if !file.IsDir() {
			fmt.Println("<", file.Name(), ">")
			ProcessFile(filepath.Join(".", "raw", file.Name()))
			fmt.Println("----------------------")
		}
	}
}

func ProcessMainDraw(wikiText string) {
	content := wiki.ExtractSection(wikiText, "Main draw")
	scanner := bufio.NewScanner(strings.NewReader(content))
	row := 1
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("Row #%v — ©%v®\n", row, line)
		if strings.Contains(line, "flagathlete") || strings.Contains(line, "flagicon") {
			params := wiki.ExtractParams(line)
			fmt.Println(strings.Join(params, "\n"))
		}
		row++
	}
}

func main() {
	markup, err := client.GetContent("2023 World Snooker Championship")
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		return
	}
	ProcessMainDraw(markup)
	fmt.Println(os.Args)
}
