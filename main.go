package main

import (
	"DEcrypt/cli"
	"DEcrypt/script"
	"DEcrypt/tool"
	"fmt"
	"path/filepath"
	"time"
)

func main() {
	// Отримуємо конфігурацію з cli.go
	cfg := cli.Cli()

	if cfg.Input == "" {
		fmt.Println("Помилка: не вказано вхідний файл.")
		fmt.Println("Використання: -i <файл> [-o <вихід>] або просто перетягніть файл на програму.")
		return
	}

	fmt.Printf("Вхідний файл: %s\n", cfg.Input)
	//костиль
	cfg.Output, _ = filepath.Abs(cfg.Input)
	cfg.Output = filepath.Dir(cfg.Output)
	tool.Log(cfg.Output)

	if cfg.Output != "" {
		fmt.Printf("Вихідний шлях: %s\n", cfg.Output)
	} else {
		fmt.Println("Вихідний шлях не вказано (буде використано значення за замовчуванням).")
	}
	// провірка формата
	if filepath.Ext(cfg.Input) != ".py" {
		println("недопустимий формат дайте файл .py")

	} else {

		defer script.Clear(cfg.Input, cfg.Output)
		script.Pac(cfg.Input, cfg.Output)
	}
	// Щоб вікно не закрилося відразу при Drag & Drop у Windows
	fmt.Println("\n3 сек до виходу...")
	time.Sleep(300000 * time.Second)
}
