package tool

import (
	"fmt"
	"os"
)

func Log(m string, args ...error) { println(m, args) }

func Err(e error) { fmt.Printf("%v\n", e); var input string; fmt.Scanln(&input); os.Exit(0) }
