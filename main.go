package main

import (
	"DEcrypt/cli"
	"DEcrypt/config"
	"DEcrypt/script"
	"DEcrypt/tool"
	"fmt"
	"time"
)

func main() {
	tool.Initlog(true)

	c := config.New()
	cli.Cli(c)

	tool.Log("Вхідний файл: %s\n", c.Input)

	tool.Log("out: %s\n", c.Output)

	defer script.Clear(c)

	script.Pac(c)

	script.Clear(c)
	// Щоб вікно не закрилося відразу при Drag & Drop у Windows
	fmt.Println("\n3 сек до виходу...")
	time.Sleep(3 * time.Second)
}
