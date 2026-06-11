package tool

import (
	"fmt"
	"os"
)

var logs bool

func Initlog(stat bool) { logs = stat }

func Log(m string, args ...any) {
	if !logs {
		return
	}
	if len(args) > 0 {
		println(m, fmt.Sprint(args))
	}
	println(m)
}

func Err(e error) { fmt.Printf("%v\n", e); var input string; fmt.Scanln(&input); os.Exit(0) }
