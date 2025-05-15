package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

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

func main() {
	markup, err := getWikiMarkup("2025 World Snooker Championship")
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		return
	}
	fmt.Println(markup)
}