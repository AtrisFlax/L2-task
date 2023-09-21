package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
)

func main() {
	var after int
	var before int
	var context int
	var count *bool
	var ignoreCase *bool
	var exclude *bool
	var fixed *bool
	var num *bool

	flag.IntVar(&after, "A", 0, "Число строк после найденной")
	flag.IntVar(&before, "B", 0, "Число строк до найденной")
	flag.IntVar(&context, "C", 0, "Число строк до и после найденной")
	count = flag.Bool("c", false, "Количество строк")
	ignoreCase = flag.Bool("i", false, "Игнорировать регистр")
	exclude = flag.Bool("v", false, "Вместо совпадения, совпадения исключать")
	fixed = flag.Bool("F", false, "Точное совпадение со строкой а не паттерн")
	num = flag.Bool("n", false, "Печатать номер строки ")

	flag.Parse()

	pattern := flag.Arg(0)

	regExp := makeRegExp(pattern, *ignoreCase)

	lines := openArgFile()

	gp := Create(lines, regExp, pattern, after, before, context, count, ignoreCase, exclude, fixed, num)

	gp.Process()
}

func openArgFile() []string {
	var reader io.Reader
	if flag.Arg(1) != "" {
		f, err := os.Open(flag.Arg(1))
		if err != nil {
			fmt.Printf("grep: %s\n", err)
			os.Exit(1)
		}
		defer func(f *os.File) {
			err := f.Close()
			if err != nil {

			}
		}(f)
		reader = f
	}
	s := bufio.NewScanner(reader)
	result := make([]string, 0)
	for s.Scan() {
		result = append(result, s.Text())
	}
	return result
}

func makeRegExp(pattern string, ignoreCase bool) *regexp.Regexp {
	var regExp *regexp.Regexp
	if ignoreCase {
		caseInsensitivePrefix := `(?i)`
		regExp = regexp.MustCompile(caseInsensitivePrefix + pattern)
	} else {
		regExp = regexp.MustCompile(pattern)
	}
	return regExp
}
