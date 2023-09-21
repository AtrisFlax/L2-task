package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {

	field := flag.Int("f", -1, "")
	delimiter := flag.String("d", " ", "")
	separated := flag.Bool("s", false, "")
	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)
	var inputStrings []string

	for scanner.Scan() {
		input := scanner.Text()
		if input == "" {
			break
		}
		inputStrings = append(inputStrings, strings.TrimSpace(input))
	}

	result := cut(inputStrings, *field, *delimiter, *separated)

	fmt.Println(result)
}

func cut(lines []string, field int, delimiter string, separated bool) string {
	var cutLines []string
	for _, row := range lines {
		fields := strings.Split(row, delimiter)
		if separated && len(fields) <= 1 {
			continue
		}
		if len(fields) > field {
			cutLines = append(cutLines, fields[field])
		}
	}
	return strings.Join(cutLines, "\n")
}
