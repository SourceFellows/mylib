package mylib

import "regexp"

var compile = regexp.MustCompile("\\d")

func HideDigits(text string) string {

	return compile.ReplaceAllLiteralString(text, "X")

}
