package reg

import (
	"regexp"
)

var replacePattern = regexp.MustCompile(
	`{{.+?}}`,
)

var pattern = regexp.MustCompile(
	`>\s*(\w.*?)\s*<`,
)

func FindMatchesInFile(file []byte) []string {
	text := string(file)
	text = replacePattern.ReplaceAllString(text, "")
	match := pattern.FindAllStringSubmatch(text, -1)
	result := make([]string, len(match))
	for i := range match {
		result[i] = match[i][1]
	}
	return result
}
