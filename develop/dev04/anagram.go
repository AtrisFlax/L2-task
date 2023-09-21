package anagram

import (
	"sort"
	"strings"
)

func anagram(input []string) map[string][]string {
	uniqueLowerCaseWords := uniqueLowerCaseStringSlice(input)

	anagramMap := make(map[string][]string)
	for _, word := range uniqueLowerCaseWords {
		sortedWord := wordWithSortedLetters(word)
		anagrams, ok := anagramMap[sortedWord]
		if ok {
			anagramMap[sortedWord] = append(anagrams, word)
		} else {
			var wordAnagrams []string
			anagramMap[sortedWord] = append(wordAnagrams, word)
		}
	}

	result := make(map[string][]string)
	for key, words := range anagramMap {
		if len(words) <= 1 {
			continue
		} else {
			result[anagramMap[key][0]] = words[1:]
		}
	}
	return result
}

func wordWithSortedLetters(lowerCaseWord string) string {
	letters := []rune(lowerCaseWord)
	sort.Slice(letters, func(i, j int) bool {
		return letters[i] < letters[j]
	})
	return string(letters)
}

func uniqueLowerCaseStringSlice(slice []string) []string {
	uniqueMap := make(map[string]bool)
	var uniqueSlice []string

	for _, str := range slice {
		lowerCaseStr := strings.ToLower(str)
		if !uniqueMap[lowerCaseStr] {
			uniqueMap[lowerCaseStr] = true
			uniqueSlice = append(uniqueSlice, lowerCaseStr)
		}
	}
	return uniqueSlice
}
