package client


func GetContent(title string) (string, error) {
        params := url.Values{}
        params.Add("action", "query")
        params.Add("prop", "revisions")
        params.Add("rvprop", "content")
        params.Add("rvslots", "main")
        params.Add("formatversion", "2")
        params.Add("format", "json")
        params.Add("titles", title)

        apiURL := "https://en.wikipedia.org/w/api.php?" + params.Encode()

        resp, err := http.Get(apiURL)
        if err != nil {
                return "", err
        }
        defer resp.Body.Close()

        body, err := ioutil.ReadAll(resp.Body)
        if err != nil {
                return "", err
        }

        var result map[string]interface{}
        if err := json.Unmarshal(body, &result); err != nil {
                return "", err
        }

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