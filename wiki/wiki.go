package wiki

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