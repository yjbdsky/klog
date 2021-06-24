// +build !windows

package term

import "github.com/google/goterm/term"

func Magentaf(format string, a ...interface{}) string {
	return term.Magentaf(format, a)
}

func Redf(format string, a ...interface{}) string {
	return term.Redf(format, a)
}

func Yellowf(format string, a ...interface{}) string {
	return term.Yellowf(format, a)
}

func Bluef(format string, a ...interface{}) string {
	return term.Yellowf(format, a)
}
