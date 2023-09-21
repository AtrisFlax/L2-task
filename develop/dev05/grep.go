package main

import (
	"fmt"
	"regexp"
	"strings"
)

type grepParams struct {
	lines   []string
	re      *regexp.Regexp
	pattern string
	//flags
	after      int
	before     int
	context    int
	count      *bool
	ignoreCase *bool
	exclude    *bool
	fixed      *bool
	num        *bool
}

func Create(lines []string, re *regexp.Regexp, pattern string,
	after, before, context int,
	count, ignoreCase, exclude, fixed, num *bool) *grepParams {
	return &grepParams{
		lines:   lines,
		re:      re,
		pattern: pattern,
		//flags
		after:      after,
		before:     before,
		context:    context,
		count:      count,
		ignoreCase: ignoreCase,
		exclude:    exclude,
		fixed:      fixed,
		num:        num,
	}
}

func (gp grepParams) Process() {
	if *gp.count {
		matchLines := countMatchLines(gp)
		fmt.Println(matchLines)
	} else {
		for lineNum := 0; lineNum < len(gp.lines); lineNum++ {
			valid := gp.isMatch(lineNum)
			isInverted := valid != *gp.exclude
			if isInverted {
				//numerate lines
				if *gp.num {
					fmt.Printf("%d:", lineNum+1)
				}
				gp.linesBefore(lineNum, gp.before)
				fmt.Println(gp.lines[lineNum])
				gp.linesAfter(lineNum, gp.after)
				gp.ContextString(lineNum)
			}
		}
	}
}

func countMatchLines(gp grepParams) int {
	matchLinesCounter := 0
	for lineNum := 0; lineNum < len(gp.lines); lineNum++ {
		valid := gp.isMatch(lineNum)
		if valid != *gp.exclude {
			matchLinesCounter++
		}
	}

	return matchLinesCounter
}

func (gp grepParams) isMatch(lineNum int) bool {
	//if fixed "-F" use exact string match
	if *gp.fixed {
		if *gp.ignoreCase {
			//ignore case
			lineLowCase := strings.ToLower(gp.lines[lineNum])
			patternLowCase := strings.ToLower(gp.pattern)
			return lineLowCase == patternLowCase
		} else {
			//don't ignore case
			line := gp.lines[lineNum]
			pattern := gp.pattern
			return strings.Contains(line, pattern)
		}
	}
	return gp.re.MatchString(gp.lines[lineNum])
}

func (gp grepParams) linesBefore(lineNum int, before int) {
	if before == 0 {
		return
	}
	first := 0
	if lineNum-before > 0 {
		first = lineNum - before
	}
	for j := first; j < lineNum; j++ {
		fmt.Println(gp.lines[j])
	}
}

func (gp grepParams) linesAfter(lineNum int, param int) {
	if param == 0 {
		return
	}
	ending := len(gp.lines) - 1
	if lineNum+param < len(gp.lines)-1 {
		ending = lineNum + param
	}
	for j := lineNum + 1; j <= ending; j++ {
		fmt.Println(gp.lines[j])
	}
}

func (gp grepParams) ContextString(lineNum int) {
	if gp.context != 0 {
		gp.linesBefore(lineNum, gp.context)
		fmt.Println(gp.lines[lineNum])
		gp.linesAfter(lineNum, gp.context)
		fmt.Println()
	}
}
