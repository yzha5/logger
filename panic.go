package yl

import "fmt"

func PanicInfo(s string) {
	panic(fmt.Sprintf("\033[36m%s\033[0m\n", s))
}
func PanicSuccess(s string) {
	panic(fmt.Sprintf("\033[32m%s\033[0m\n", s))
}
func PanicWarn(s string) {
	panic(fmt.Sprintf("\033[33m%s\033[0m\n", s))
}
func PanicError(s string) {
	panic(fmt.Sprintf("\033[31m%s\033[0m\n", s))
}

func PanicInfof(format string, a ...any) {
	panic(fmt.Sprintf(fmt.Sprintf("\033[36m%s\033[0m\n", format), a...))
}
func PanicSuccessf(format string, a ...any) {
	panic(fmt.Sprintf(fmt.Sprintf("\033[32m%s\033[0m\n", format), a...))
}
func PanicWarnf(format string, a ...any) {
	panic(fmt.Sprintf(fmt.Sprintf("\033[33m%s\033[0m\n", format), a...))
}
func PanicErrorf(format string, a ...any) {
	panic(fmt.Sprintf(fmt.Sprintf("\033[31m%s\033[0m\n", format), a...))
}

func PanicInfoBg(s string) {
	panic(fmt.Sprintf("\033[46;30m%s\033[0m\n", s))
}
func PanicSuccessBg(s string) {
	panic(fmt.Sprintf("\033[42;30m%s\033[0m\n", s))
}
func PanicWarnBg(s string) {
	panic(fmt.Sprintf("\033[43;30m%s\033[0m\n", s))
}
func PanicErrorBg(s string) {
	panic(fmt.Sprintf("\033[41;30m%s\033[0m\n", s))
}

func PanicInfoBgf(format string, a ...any) {
	panic(fmt.Sprintf(fmt.Sprintf("\033[46;30m%s\033[0m\n", format), a...))
}
func PanicSuccessBgf(format string, a ...any) {
	panic(fmt.Sprintf(fmt.Sprintf("\033[42;30m%s\033[0m\n", format), a...))
}
func PanicWarnBgf(format string, a ...any) {
	panic(fmt.Sprintf(fmt.Sprintf("\033[43;30m%s\033[0m\n", format), a...))
}
func PanicErrorBgf(format string, a ...any) {
	panic(fmt.Sprintf(fmt.Sprintf("\033[41;30m%s\033[0m\n", format), a...))
}
