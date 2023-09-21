package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type digitAndLine struct {
	firstDigitFromStart *int
	line                *string
}

type targetWordWithLine struct {
	word *string
	line *string
}

func main() {
	isReverseSortOrder := flag.Bool("r", false, "сортировать в обратном порядке")
	isColumnSort := flag.String("k", "1", "указание колонки для сортировки (слова в строке могут выступать в качестве колонок, по умолчанию разделитель — пробел")
	isNumeralSort := flag.Bool("n", false, "сортировать по числовому значению")
	isUnique := flag.Bool("u", false, "не выводить повторяющиеся строки")
	flag.Parse()

	if len(os.Args) < 2 {
		_, _ = fmt.Fprintf(os.Stderr, "Missing parameter, provide file name!")
		os.Exit(1)
	}

	input, err := os.ReadFile(flag.Arg(0))
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Can't open file")
		os.Exit(1)
	}

	lines := strings.Split(string(input), "\n")

	if *isUnique {
		lines = unique(lines)
	}

	switch {
	case isColumnSort != nil && isNumeralSort != nil:
		lines = columnAndNumericSort(lines, *isColumnSort)
	case isColumnSort == nil && isNumeralSort != nil:
		lines = numericSort(lines)
	case isColumnSort != nil && isNumeralSort == nil:
		lines = columnSort(lines, *isColumnSort)
	case isColumnSort == nil && isNumeralSort == nil:
		basicSort(lines)
	}

	if *isReverseSortOrder {
		reverse(lines)
	}

	for _, line := range lines {
		fmt.Println(line)
	}
}

func basicSort(lines []string) {
	isLessFunc := func(i int, j int) bool {
		return lines[i] < lines[j]
	}
	sort.Slice(lines, isLessFunc)
}

func reverse(lines []string) {
	for i, j := 0, len(lines)-1; i < j; i, j = i+1, j-1 {
		lines[i], lines[j] = lines[j], lines[i]
	}
}

func unique(lines []string) []string {
	h := map[string]struct{}{}
	for _, v := range lines {
		h[v] = struct{}{}
	}

	unique := make([]string, 0, len(h))
	for k := range h {
		unique = append(unique, k)
	}
	return unique
}

func numericSort(lines []string) []string {
	var numerics, nonNumerics []digitAndLine
	numerics, nonNumerics = getNumericAndNonNumericLines(lines, numerics, nonNumerics)

	sort.Slice(numerics, func(i, j int) bool {
		return *numerics[i].firstDigitFromStart < *numerics[j].firstDigitFromStart
	})
	sort.Slice(nonNumerics, func(i, j int) bool {
		return *nonNumerics[i].line < *nonNumerics[j].line
	})

	result := make([]string, 0, len(lines))
	for i := 0; i < len(nonNumerics); i++ {
		result = append(result, *nonNumerics[i].line)
	}
	for i := 0; i < len(numerics); i++ {
		result = append(result, *numerics[i].line)
	}
	return result
}

func columnSort(lines []string, keyDef string) []string {
	startNum := getStartNum(keyDef)
	startIndex := startNum - 1

	targets := getColumnTargets(lines, startIndex)

	sort.Slice(targets, func(i, j int) bool {
		return *targets[i].word < *targets[j].word
	})

	result := make([]string, 0, len(lines))
	for i := 0; i < len(targets); i++ {
		result = append(result, *targets[i].line)
	}
	return result
}

func getNumericAndNonNumericLines(lines []string, numerics []digitAndLine, nonNumerics []digitAndLine) ([]digitAndLine, []digitAndLine) {
	for i, line := range lines {
		if firstDigitStr := regexp.MustCompile(`^\d+`).FindString(line); firstDigitStr != "" {
			digit, _ := strconv.Atoi(firstDigitStr)
			numerics = append(numerics, digitAndLine{
				firstDigitFromStart: &digit,
				line:                &lines[i],
			})
		} else {
			nonNumerics = append(nonNumerics, digitAndLine{
				firstDigitFromStart: nil,
				line:                &lines[i],
			})
		}
	}
	return numerics, nonNumerics
}

func getColumnTargets(lines []string, startIndex int) []targetWordWithLine {
	targets := make([]targetWordWithLine, 0, len(lines))
	for i := 0; i < len(lines); i++ {
		words := strings.Split(lines[i], " ")
		var targetWord string
		if startIndex < len(words) {
			targetWord = words[startIndex]
		} else {
			targetWord = ""
		}
		targets = append(targets, targetWordWithLine{&targetWord, &lines[i]})
	}
	return targets
}

func getStartNum(keyDef string) int {
	startNum, err := strconv.Atoi(keyDef)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "invalid number at field start: '%s' not a number", keyDef)
		os.Exit(1)
	}
	if startNum <= 0 {
		_, _ = fmt.Fprintf(os.Stderr, "invalid number at field start: invalid count at start of '%s'", keyDef)
		os.Exit(1)
	}
	return startNum
}

func columnAndNumericSort(lines []string, keyDef string) []string {
	startNum := getStartNum(keyDef)
	startIndex := startNum - 1

	numerics, nonNumerics := getNumericAndNonNumericWords(lines, startIndex)

	sort.Slice(numerics, func(i, j int) bool {
		return *numerics[i].firstDigitFromStart < *numerics[j].firstDigitFromStart
	})

	sort.Slice(nonNumerics, func(i, j int) bool {
		return *nonNumerics[i].line < *nonNumerics[j].line
	})

	result := make([]string, 0, len(lines))
	for i := 0; i < len(nonNumerics); i++ {
		result = append(result, *nonNumerics[i].line)
	}
	for i := 0; i < len(numerics); i++ {
		result = append(result, *numerics[i].line)
	}
	return result
}

func getNumericAndNonNumericWords(lines []string, startIndex int) ([]digitAndLine, []digitAndLine) {
	var numerics, nonNumerics []digitAndLine
	for i := 0; i < len(lines); i++ {
		words := strings.Split(lines[i], " ")
		if startIndex < len(words) {
			targetWord := words[startIndex]
			if firstDigitOfWord := regexp.MustCompile(`^\d+`).FindString(targetWord); firstDigitOfWord != "" {
				digit, _ := strconv.Atoi(firstDigitOfWord)
				numerics = append(numerics, digitAndLine{
					firstDigitFromStart: &digit,
					line:                &lines[i],
				})
			} else {
				nonNumerics = append(nonNumerics, digitAndLine{
					firstDigitFromStart: nil,
					line:                &lines[i],
				})
			}
		} else {
			nonNumerics = append(nonNumerics, digitAndLine{
				firstDigitFromStart: nil,
				line:                &lines[i],
			})
		}
	}
	return numerics, nonNumerics
}
