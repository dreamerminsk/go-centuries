package wiki

import (
	"bufio"
	"strings"
)

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

func ExtractParams(wikiText string) []string {
	params := []string{}
	openCount := 0
	var param strings.Builder
	for i := 0; i < len(wikiText)-1; i++ {
		if wikiText[i] == '|' && openCount == 0 {
			params = append(params, param.String())
			param.Reset()
		} else if wikiText[i] == '{' {
			param.WriteByte(wikiText[i])
			openCount++
		} else if wikiText[i] == '[' {
			param.WriteByte(wikiText[i])
			openCount++
		} else if wikiText[i] == '}' {
			param.WriteByte(wikiText[i])
			openCount--
		} else if wikiText[i] == ']' {
			param.WriteByte(wikiText[i])
			openCount--
		} else {
			param.WriteByte(wikiText[i])
		}
	}
 if param.Len() > 0 {
    params = append(params, param.String())
  }
	return params
}
