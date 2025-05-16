package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
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

func getWikiMarkup(title string) (string, error) {
	// Устанавливаем параметры запроса
	params := url.Values{}
	params.Add("action", "query")
	params.Add("prop", "revisions")
	params.Add("rvprop", "content")
	params.Add("rvslots", "main")
	params.Add("formatversion", "2")
	params.Add("format", "json")
	params.Add("titles", title)

	// Формируем URL запроса
	apiURL := "https://en.wikipedia.org/w/api.php?" + params.Encode()

	// Выполняем HTTP GET запрос
	resp, err := http.Get(apiURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Читаем тело ответа
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Парсим JSON ответ
	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return "", err
	}

	// Извлекаем разметку из структуры ответа
	pages := result["query"].(map[string]interface{})["pages"].([]interface{})
	page := pages[0].(map[string]interface{})
	if _, ok := page["missing"]; ok {
		return "", fmt.Errorf("страница не найдена")
	}
	revisions := page["revisions"].([]interface{})
	revision := revisions[0].(map[string]interface{})
	slots := revision["slots"].(map[string]interface{})
	main := slots["main"].(map[string]interface{})
	content := main["content"].(string)

	return content, nil
}

func ExtractFullSection(wikiText, sectionName string) string {
	var sectionContent []string
	inSection := false
	sectionLevel := 0
	sectionName = strings.TrimSpace(strings.ToLower(sectionName))

	scanner := bufio.NewScanner(strings.NewReader(wikiText))
	for scanner.Scan() {
		line := scanner.Text()
		trimmedLine := strings.TrimSpace(line)
		if strings.HasPrefix(trimmedLine, "=") && strings.HasSuffix(trimmedLine, "=") {
			headerText := strings.Trim(trimmedLine, "=")
			headerText = strings.TrimSpace(headerText)
			levelCount := len(trimmedLine) - len(strings.TrimLeft(trimmedLine, "="))

			if inSection && levelCount <= sectionLevel {
				break
			}

			if strings.ToLower(headerText) == sectionName {
				inSection = true
				sectionLevel = levelCount
				continue
			}
		}

		if inSection {
			sectionContent = append(sectionContent, line)
		}
	}
	return strings.TrimSpace(strings.Join(sectionContent, "\n"))
}

func ExtractSection(wikiText, sectionName string) string {
	var sectionContent []string
	inSection := false
	sectionName = strings.TrimSpace(strings.ToLower(sectionName))

	scanner := bufio.NewScanner(strings.NewReader(wikiText))
	for scanner.Scan() {
		line := scanner.Text()
		trimmedLine := strings.TrimSpace(line)
		if strings.HasPrefix(trimmedLine, "=") && strings.HasSuffix(trimmedLine, "=") {
			headerText := strings.Trim(trimmedLine, "=")
			headerText = strings.TrimSpace(headerText)

			if inSection {
				break
			}

			if strings.ToLower(headerText) == sectionName {
				inSection = true
			}
		}

		if inSection {
			sectionContent = append(sectionContent, line)
		}
	}
	return strings.TrimSpace(strings.Join(sectionContent, "\n"))
}

func ProcessMainDraw(wikiText string) {
	content := ExtractSection(wikiText, "Main draw")
	scanner := bufio.NewScanner(strings.NewReader(content))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "flagathlete") {
			fmt.Println(line)
		}
	}
}

func main() {
	markup, err := getWikiMarkup("1985 World Snooker Championship")
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		return
	}
	ProcessMainDraw(markup)
	fmt.Println(os.Args)
}
