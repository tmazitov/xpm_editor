package xpm

import (
	"fmt"
	"regexp"
	"strings"
)

type xpmColor struct {
	character rune
	color     string
}

func NewXpmColor(row string) *xpmColor {

	row = strings.ReplaceAll(row, ",", "")
	row = strings.ReplaceAll(row, "\"", "")
	return &xpmColor{
		character: rune(row[0]),
		color:     row[4:],
	}
}

func (x *xpmColor) ToString() string {
	return fmt.Sprintf("\"%c c %s\",", x.character, x.color)
}

func isColor(line string) bool {
	pattern := `.[ \t]c #[0-9A-Fa-f]{6}$`

	// Compile the regular expression
	re, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return false
	}

	line = strings.ReplaceAll(line, ",", "")
	line = strings.ReplaceAll(line, "\"", "")
	// Test a string against the regular expression
	return re.MatchString(line)
}
