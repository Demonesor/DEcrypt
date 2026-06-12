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
	start := time.Now()
	tool.Initlog(true)

	c := config.New()
	cli.Cli(c)

	tool.Log("Вхідний файл: %s\n", c.Paths.AbsInput)

	tool.Log("out: %s\n", c.Paths.AbsOutput)

	defer script.Clear(c)

	script.Start(c)
	script.Clear(c)

	// Щоб вікно не закрилося відразу при Drag & Drop у Windows
	duration := time.Since(start)

	fmt.Printf("Час виконання: %v\n", duration)
	fmt.Println("\n3 сек до виходу...")
	time.Sleep(3 * time.Second)
}
