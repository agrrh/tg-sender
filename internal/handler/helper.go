package handler

import (
	"strings"
)

func fmtTelegram(input string) string {
	return strings.NewReplacer(
		"(", "\\(", // ()_-. are reserved by telegram.
		")", "\\)",
		"_", "\\_",
		"*", "\\*",
		".", "\\.",
		"-", "\\-",
		"!", "\\!",
		"#", "\\#",
		"+", "\\+",
		"=", "\\=",
		"{", "\\{",
		"}", "\\}",
		"<", "\\<",
		">", "\\>",
		"~", "\\~",
	).Replace(input)
}
