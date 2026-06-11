package tool

import (
	"DEcrypt/config"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func Tempdir(c *config.Config) {
	c.TempDir = os.TempDir()

	// 2. Рахуємо хеш файлу
	hash, err := getFileHash(c.Input)
	if err != nil {
		fmt.Printf("Помилка при читанні файлу: %v\n", err)
		return
	}

	// 3. Формуємо ім'я папки
	folderName := filepath.Base(c.Input) + hash

	// 4. Склеюємо ПОВНИЙ абсолютний шлях до нашої робочої папки всередині Temp
	// Результат: C:\Users\...\AppData\Local\Temp\test.py901b841d...
	c.Temp = filepath.Join(c.TempDir, folderName)

	// 5. Створюємо цю папку (MkdirAll не панікує, якщо папка вже існує)
	err = os.MkdirAll(c.Temp, 0755)
	if err != nil {
		fmt.Printf("Помилка створення папки в Temp: %v\n", err)
		return
	}

	// 6. Повертаємо ПОВНИЙ шлях, щоб інші функції знали, куди копіювати і де шукати python.exe
	return
}

func CopyFile(src, dst string) error {
	// Відкриваємо вхідний файл
	Log(src)
	Log(dst)
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close() // Гарантуємо закриття після завершення

	// Створюємо вихідний файл
	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close() // Гарантуємо закриття

	// Переливаємо дані з входу на вихід
	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return err
	}

	return nil
}
