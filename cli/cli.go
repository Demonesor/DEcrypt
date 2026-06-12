package cli // або який у тебе тут пакет

import (
	"DEcrypt/config"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

// Передаємо ВКАЗІВНИК на головний конфіг, щоб змінювати оригінал
func Cli(c *config.Config) {

	// Прив'язуємо прапорці до полів нашого конфігу
	flag.StringVar(&c.Paths.Input, "i", "", "Шлях до вхідного файлу .py")
	flag.StringVar(&c.Paths.Output, "o", "", "Шлях до вихідного файлу або папки")

	// Парсимо аргументи консолі
	flag.Parse()

	// Обробка Drag & Drop (якщо -i не вказано)
	if c.Paths.Input == "" {
		args := flag.Args()
		if len(args) > 0 {
			c.Paths.Input = args[0]
		}
	}

	// 1. Жорстка перевірка: чи взагалі є вхідний файл?
	if c.Paths.Input == "" {
		fmt.Println("❌ Помилка: не вказано вхідний файл.")
		fmt.Println("💡 Використання: DECrypt.exe -i <файл.py> [-o <вихід>] або просто перетягніть файл на програму.")
		os.Exit(1) // Зупиняємо програму з кодом помилки
	}

	// 2. Перевірка формату (.py)
	if filepath.Ext(c.Paths.Input) != ".py" {
		fmt.Println("❌ Помилка: недопустимий формат. Програма приймає тільки файли .py")
		os.Exit(1) // Зупиняємо програму
	}

	// 3. Тепер, коли ми точно маємо Input, фіксуємо його абсолютний шлях
	absInput, err := filepath.Abs(c.Paths.Input)
	if err != nil {
		fmt.Println("❌ Помилка обробки шляху вхідного файлу:", err)
		os.Exit(1)
	}
	c.Paths.AbsInput = absInput

	// 4. Логіка для Output (якщо користувач його не вказав)
	if c.Paths.Output == "" {
		// Якщо вихідна папка не вказана, зберігаємо екзешник у тій самій папці, де лежить скрипт
		c.Paths.Output = filepath.Dir(c.Paths.AbsInput)

		fmt.Println("⚠️ Вихідний шлях не вказано. Файл буде збережено поруч із скриптом:", c.Paths.Output)
	} else {
		// Якщо вказав, отримуємо абсолютний шлях
		c.Paths.AbsOutput, _ = filepath.Abs(c.Paths.Output)
		fmt.Printf("✅ Вихідний шлях встановлено: %s\n", c.Paths.AbsOutput)
	}
}
