// 2
package main

import (
	"fmt"
	"sort"
	"strings"
)

func getTopWords(wordMap map[string]int, n int) []string {
	keys := make([]string, 0, len(wordMap))
	for k := range wordMap {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return wordMap[keys[i]] > wordMap[keys[j]]
	})
	if n > len(keys) {
		n = len(keys)
	}
	return keys[:n]
}

func AnalyzeText(text string) {
	text = strings.ReplaceAll(text, ",", " ")
	text = strings.ReplaceAll(text, ".", " ")
	text = strings.ReplaceAll(text, "!", " ")
	text = strings.ReplaceAll(text, "?", " ")
	rawWords := strings.Fields(text)
	words := make([]string, 0, len(rawWords))
	for _, w := range rawWords {
		words = append(words, strings.ToLower(w))
	}
	wordCount := make(map[string]int)
	for _, w := range words {
		wordCount[w]++
	}
	totalWords := len(words)
	uniqueWords := len(wordCount)
	maxWord := ""
	maxCount := 0
	for w, c := range wordCount {
		if c > maxCount {
			maxCount = c
			maxWord = w
		}
	}
	top5 := getTopWords(wordCount, 5)
	fmt.Printf("Количество слов: %d\n", totalWords)
	fmt.Printf("Количество уникальных слов: %d\n", uniqueWords)
	fmt.Printf("Самое часто встречающееся слово: \"%s\" (встречается %d раз)\n", maxWord, maxCount)
	fmt.Println("Топ-5 самых часто встречающихся слов:")
	for _, w := range top5 {
		fmt.Printf("\"%s\": %d раз\n", w, wordCount[w])
	}
}
