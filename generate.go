package main

import (
	"os"
	"strings"
)

func GenerateArt(input string, banner string) string {
	input = strings.ReplaceAll(input, "\\n", "\n")
	input = strings.ReplaceAll(input, "\r\n", "\n")
	filepath := "banners/" + banner + ".txt"

	data, err := os.ReadFile(filepath)

	if err != nil {
		return "error reading file"
	}

	characters := strings.ReplaceAll(string(data), "\r\n", "\n")

	lines := strings.Split(characters, "\n")

	words := strings.Split(input, "\n")

	var result strings.Builder

	for _, word := range words {
		if word == "" {
			result.WriteString("\n")
			continue
		}
		for row := 0; row < 8; row++ {
			for _, char := range word {
				index := int(char-32)*9 + row

				if index < len(lines) {
					result.WriteString(lines[index])

				}
				//result.WriteString(lines[index])
			}
			result.WriteString("\n")
		}
	}

	return result.String()
}
