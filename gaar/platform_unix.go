// +build !windows

package gaar

import (
	"strings"
)

const EnvVarSetKeyword = "export"

func sanitizeMessage(msg string) string {
	quoteChar := "'"
	msg = strings.Replace(msg, quoteChar, "'\"\\'\"'", -1)
	msg = strings.TrimRight(msg, "\\")
	return quoteChar + msg + quoteChar
}
