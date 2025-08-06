package stringUtil

import "strings"

func AddBothSidesPadding(s string) string {
	return " " + s + " "
}

func AddRightSidePadding(s string, maxWidth int) string {
	remainingWidth := maxWidth - len(s)

	if remainingWidth < 0 {
		return s
	}

	return s + strings.Repeat(" ", remainingWidth)
}

func AddPipes(sList []string) string {
	return "|" + strings.Join(sList, "|") + "|"
}

func HeaderSeparator(headerRowLength int) string {
	return strings.Repeat("-", headerRowLength)
}
