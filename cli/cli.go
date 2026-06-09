package cli

import (
	"flag"
	"os"
)

// Config зберігає параметри командного рядка
type Config struct {
	Input  string
	Output string
}

func Cli() Config {
	var cfg Config

	// Визначаємо базові прапорці
	// -i: вхідний файл
	// -o: вихідний файл або папка
	flag.StringVar(&cfg.Input, "i", "", "Шлях до вхідного файлу")
	flag.StringVar(&cfg.Output, "o", "", "Шлях до вихідного файлу/папки")

	// Парсимо аргументи
	flag.Parse()

	// Обробка Drag & Drop:
	// Якщо прапорець -i не вказано, але є "вільні" аргументи (flag.Args),
	// беремо перший з них як вхідний файл.
	if cfg.Input == "" {
		args := flag.Args()
		if len(args) > 0 {
			cfg.Input = args[0]
		} else if len(os.Args) > 1 && os.Args[1][0] != '-' {
			// Якщо прапорців взагалі немає, але є аргумент
			cfg.Input = os.Args[1]
		} else {
			return Config{"", ""}
		}
	}

	return cfg
}
