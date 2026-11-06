package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"strings"
)

//go:embed content.json
var contentFS embed.FS

type Section struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}
type Content struct {
	Name     string    `json:"name"`
	Sections []Section `json:"sections"`
}

func loadContent() (Content, error) {
	b, err := contentFS.ReadFile("content.json")
	if err != nil {
		return Content{}, err
	}
	return parseContent(b)
}
func parseContent(b []byte) (Content, error) {
	var c Content
	if err := json.Unmarshal(b, &c); err != nil {
		return c, err
	}
	if c.Name == "" || len(c.Sections) != 8 {
		return c, fmt.Errorf("content requires a name and eight sections")
	}
	for _, s := range c.Sections {
		if strings.TrimSpace(s.Title) == "" || strings.TrimSpace(s.Body) == "" || len(s.Body) > 16000 {
			return c, fmt.Errorf("invalid section")
		}
		for _, r := range s.Body + s.Title {
			if r < 32 && r != '\n' {
				return c, fmt.Errorf("control character in content")
			}
		}
	}
	return c, nil
}
