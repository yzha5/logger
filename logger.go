package yl

import (
	"fmt"
	"log"
)

func LogInfo(s string) {
	log.Printf("\033[36m%s\033[0m\n", s)
}
func LogSuccess(s string) {
	log.Printf("\033[32m%s\033[0m\n", s)
}
func LogWarn(s string) {
	log.Printf("\033[33m%s\033[0m\n", s)
}
func LogError(s string) {
	log.Printf("\033[31m%s\033[0m\n", s)
}

func LogInfof(format string, a ...any) {
	log.Printf(fmt.Sprintf("\033[36m%s\033[0m\n", format), a...)
}
func LogSuccessf(format string, a ...any) {
	log.Printf(fmt.Sprintf("\033[32m%s\033[0m\n", format), a...)
}
func LogWarnf(format string, a ...any) {
	log.Printf(fmt.Sprintf("\033[33m%s\033[0m\n", format), a...)
}
func LogErrorf(format string, a ...any) {
	log.Printf(fmt.Sprintf("\033[31m%s\033[0m\n", format), a...)
}

func LogInfoBg(s string) {
	log.Printf("\033[46;30m%s\033[0m\n", s)
}
func LogSuccessBg(s string) {
	log.Printf("\033[42;30m%s\033[0m\n", s)
}
func LogWarnBg(s string) {
	log.Printf("\033[43;30m%s\033[0m\n", s)
}
func LogErrorBg(s string) {
	log.Printf("\033[41;30m%s\033[0m\n", s)
}

func LogInfoBgf(format string, a ...any) {
	log.Printf(fmt.Sprintf("\033[46;30m%s\033[0m\n", format), a...)
}
func LogSuccessBgf(format string, a ...any) {
	log.Printf(fmt.Sprintf("\033[42;30m%s\033[0m\n", format), a...)
}
func LogWarnBgf(format string, a ...any) {
	log.Printf(fmt.Sprintf("\033[43;30m%s\033[0m\n", format), a...)
}
func LogErrorBgf(format string, a ...any) {
	log.Printf(fmt.Sprintf("\033[41;30m%s\033[0m\n", format), a...)
}

func LogFatalInfo(s string) {
	log.Fatalf("\033[36m%s\033[0m\n", s)
}
func LogFatalSuccess(s string) {
	log.Fatalf("\033[32m%s\033[0m\n", s)
}
func LogFatalWarn(s string) {
	log.Fatalf("\033[33m%s\033[0m\n", s)
}
func LogFatalError(s string) {
	log.Fatalf("\033[31m%s\033[0m\n", s)
}

func LogFatalInfof(format string, a ...any) {
	log.Fatalf(fmt.Sprintf("\033[36m%s\033[0m\n", format), a...)
}
func LogFatalSuccessf(format string, a ...any) {
	log.Fatalf(fmt.Sprintf("\033[32m%s\033[0m\n", format), a...)
}
func LogFatalWarnf(format string, a ...any) {
	log.Fatalf(fmt.Sprintf("\033[33m%s\033[0m\n", format), a...)
}
func LogFatalErrorf(format string, a ...any) {
	log.Fatalf(fmt.Sprintf("\033[31m%s\033[0m\n", format), a...)
}

func LogFatalInfoBg(s string) {
	log.Fatalf("\033[46;30m%s\033[0m\n", s)
}
func LogFatalSuccessBg(s string) {
	log.Fatalf("\033[42;30m%s\033[0m\n", s)
}
func LogFatalWarnBg(s string) {
	log.Fatalf("\033[43;30m%s\033[0m\n", s)
}
func LogFatalErrorBg(s string) {
	log.Fatalf("\033[41;30m%s\033[0m\n", s)
}

func LogFatalInfoBgf(format string, a ...any) {
	log.Fatalf(fmt.Sprintf("\033[46;30m%s\033[0m\n", format), a...)
}
func LogFatalSuccessBgf(format string, a ...any) {
	log.Fatalf(fmt.Sprintf("\033[42;30m%s\033[0m\n", format), a...)
}
func LogFatalWarnBgf(format string, a ...any) {
	log.Fatalf(fmt.Sprintf("\033[43;30m%s\033[0m\n", format), a...)
}
func LogFatalErrorBgf(format string, a ...any) {
	log.Fatalf(fmt.Sprintf("\033[41;30m%s\033[0m\n", format), a...)
}
