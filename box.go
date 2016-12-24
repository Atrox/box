package box

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

const singleLineFormat = "%[1]s%[2]s%[3]s%[2]s%[4]s%[1]s"
const blockFormat = "%[1]s\n%s\n%[1]s\n"

type Box struct {
	HorizontalSeparator string
	VerticalSeparator   string
	SideMargin          int
}

func New() *Box {
	return &Box{
		HorizontalSeparator: "=",
		VerticalSeparator:   "|",
		SideMargin:          2,
	}
}

func (b *Box) String(lines ...string) string {
	sideMargin := strings.Repeat(" ", b.SideMargin)
	longestLine := longestLine(lines)

	// create top and bottom bars
	bar := strings.Repeat(b.HorizontalSeparator, longestLine+b.SideMargin+4)

	// create lines to print
	var texts []string
	for _, line := range lines {
		length := utf8.RuneCountInString(line)

		// use later
		var space string
		var oddSpace string

		// if current text is shorter than the longest one
		// center the text, so it looks better
		if length < longestLine {
			// difference between longest and current one
			diff := longestLine - length

			// the spaces to add on each side
			toAdd := diff / 2
			space = strings.Repeat(" ", toAdd)

			// if the difference between the longest and current one
			// is odd, we have to add one additional space before the last vertical separator
			if diff%2 != 0 {
				oddSpace = " "
			}
		}

		texts = append(texts, fmt.Sprintf(singleLineFormat, b.VerticalSeparator, space+sideMargin, line, oddSpace))
	}

	return fmt.Sprintf(blockFormat, bar, strings.Join(texts, "\n"))
}

func (b *Box) Print(lines ...string) {
	fmt.Print(b.String(lines...))
}

// Println adds a newline before and after the box
func (b *Box) Println(lines ...string) {
	fmt.Printf("\n%s\n", b.String(lines...))
}

func longestLine(lines []string) (longest int) {
	for _, line := range lines {
		length := utf8.RuneCountInString(line)
		if length > longest {
			longest = length
		}
	}
	return
}
