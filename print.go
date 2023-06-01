package yl

import (
	"fmt"
)

func PrintInfo(s string) {
	fmt.Printf("\033[36m%s\033[0m", s)
}
func PrintSuccess(s string) {
	fmt.Printf("\033[32m%s\033[0m", s)
}
func PrintWarn(s string) {
	fmt.Printf("\033[33m%s\033[0m", s)
}
func PrintError(s string) {
	fmt.Printf("\033[31m%s\033[0m", s)
}

func PrintInfof(format string, a ...any) {
	fmt.Printf(fmt.Sprintf("\033[36m%s\033[0m", format), a...)
}
func PrintSuccessf(format string, a ...any) {
	fmt.Printf(fmt.Sprintf("\033[32m%s\033[0m", format), a...)
}
func PrintWarnf(format string, a ...any) {
	fmt.Printf(fmt.Sprintf("\033[33m%s\033[0m", format), a...)
}
func PrintErrorf(format string, a ...any) {
	fmt.Printf(fmt.Sprintf("\033[31m%s\033[0m", format), a...)
}

func PrintInfoBg(s string) {
	fmt.Printf("\033[46;30m%s\033[0m", s)
}
func PrintSuccessBg(s string) {
	fmt.Printf("\033[42;30m%s\033[0m", s)
}
func PrintWarnBg(s string) {
	fmt.Printf("\033[43;30m%s\033[0m", s)
}
func PrintErrorBg(s string) {
	fmt.Printf("\033[41;30m%s\033[0m", s)
}

func PrintInfoBgf(format string, a ...any) {
	fmt.Printf(fmt.Sprintf("\033[46;30m%s\033[0m", format), a...)
}
func PrintSuccessBgf(format string, a ...any) {
	fmt.Printf(fmt.Sprintf("\033[42;30m%s\033[0m", format), a...)
}
func PrintWarnBgf(format string, a ...any) {
	fmt.Printf(fmt.Sprintf("\033[43;30m%s\033[0m", format), a...)
}
func PrintErrorBgf(format string, a ...any) {
	fmt.Printf(fmt.Sprintf("\033[41;30m%s\033[0m", format), a...)
}
