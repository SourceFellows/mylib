package mylib

import "regexp"

func HideDigits(text string) (string, error) {

	compile, err := regexp.Compile("\\d")
	if err != nil {
		return "", err
	}
	return compile.ReplaceAllLiteralString(text, "X"), nil

}
