package yl

import (
	"fmt"
)

func PrintlnInfo(s string) {
	fmt.Printf("\033[36m%s\033[0m\n", s)
}
func PrintlnSuccess(s string) {
	fmt.Printf("\033[32m%s\033[0m\n", s)
}
func PrintlnWarn(s string) {
	fmt.Printf("\033[33m%s\033[0m\n", s)
}
func PrintlnError(s string) {
	fmt.Printf("\033[31m%s\033[0m\n", s)
}

func PrintlnInfof(format string, a ...any) {
	fmt.Printf(fmt.Sprintf("\033[36m%s\033[0m\n", format), a...)
}
func PrintlnSuccessf(format string, a ...any) {
	fmt.Printf(fmt.Sprintf("\033[32m%s\033[0m\n", format), a...)
}
func PrintlnWarnf(format string, a ...any) {
	fmt.Printf(fmt.Sprintf("\033[33m%s\033[0m\n", format), a...)
}
func PrintlnErrorf(format string, a ...any) {
	fmt.Printf(fmt.Sprintf("\033[31m%s\033[0m\n", format), a...)
}

func PrintlnInfoBg(s string) {
	fmt.Printf("\033[46;30m%s\033[0m\n", s)
}
func PrintlnSuccessBg(s string) {
	fmt.Printf("\033[42;30m%s\033[0m\n", s)
}
func PrintlnWarnBg(s string) {
	fmt.Printf("\033[43;30m%s\033[0m\n", s)
}
func PrintlnErrorBg(s string) {
	fmt.Printf("\033[41;30m%s\033[0m\n", s)
}

func PrintlnInfoBgf(format string, a ...any) {
	fmt.Printf(fmt.Sprintf("\033[46;30m%s\033[0m\n", format), a...)
}
func PrintlnSuccessBgf(format string, a ...any) {
	fmt.Printf(fmt.Sprintf("\033[42;30m%s\033[0m\n", format), a...)
}
func PrintlnWarnBgf(format string, a ...any) {
	fmt.Printf(fmt.Sprintf("\033[43;30m%s\033[0m\n", format), a...)
}
func PrintlnErrorBgf(format string, a ...any) {
	fmt.Printf(fmt.Sprintf("\033[41;30m%s\033[0m\n", format), a...)
}
