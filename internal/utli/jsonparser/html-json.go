package jsonparser

import "strings"

// Only for https://hh.ru/vacancy/xxxxxx
func GetJsonFromHTML(html string) string {
	startStr := `<script type="application/ld+json">`
	startIndex := strings.Index(html, startStr)
	if startIndex == -1 {
		return ""
	}
	startIndex += len(startStr)
	stopStr := `}</script>`
	stopIndex := strings.Index(html, stopStr)
	return html[startIndex:stopIndex] + "}"
}
