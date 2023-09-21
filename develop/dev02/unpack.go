package unpack

import (
	"regexp"
	"strconv"
	"strings"
)

func DoUnpack(input string) string {
	if input == "" {
		return ""
	}
	result := strings.Builder{}
	groupsExp := regexp.MustCompile(`(\D+)(\d+)`)
	parts := groupsExp.FindAllString(input, -1)
	for _, part := range parts {
		wordRegexp := regexp.MustCompile(`(\D*)(\D)(\d+)`)
		group := wordRegexp.FindStringSubmatch(part)
		beforeLetter := group[1]
		letter := group[2]
		repeatNum, _ := strconv.Atoi(group[3])
		if beforeLetter != "" {
			result.WriteString(beforeLetter)
		}
		result.WriteString(strings.Repeat(letter, repeatNum))
	}

	endingExp := regexp.MustCompile(`(\D+)$`)
	ending := endingExp.FindString(input)
	result.WriteString(ending)
	return result.String()
}
