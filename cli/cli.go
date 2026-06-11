package cli

import (
	"DEcrypt/config"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

// Config зберігає параметри командного рядка

func Cli(c *config.Config) {

	// -i: вхідний файл
	// -o: вихідний файл або папка
	flag.StringVar(&c.Input, "i", "", "Шлях до вхідного файлу")
	flag.StringVar(&c.Output, "o", "", "Шлях до вихідного файлу/папки")

	// Парсимо аргументи
	flag.Parse()
	c.Output, _ = filepath.Abs(c.Input)

	// Обробка Drag & Drop:
	// беремо перший з них як вхідний файл.
	if c.Input == "" {
		args := flag.Args()
		if len(args) > 0 {
			c.Input = args[0]
		} else if len(os.Args) > 1 && os.Args[1][0] != '-' {
			// Якщо прапорців взагалі немає, але є аргумент
			c.Input = os.Args[1]
		}
	}

	if c.Input == "" {
		fmt.Println("Помилка: не вказано вхідний файл.")
		fmt.Println("Використання: -i <файл> [-o <вихід>] або просто перетягніть файл на програму.")
		return
	}

	if c.Output != "" {
		fmt.Printf("Вихідний шлях: %s\n", c.Output)
	} else {
		fmt.Println("Вихідний шлях не вказано (буде використано значення за замовчуванням).", c.Output)
	}

	// провірка формата
	if filepath.Ext(c.Input) != ".py" {
		println("недопустимий формат дайте файл .py")

	}

}
