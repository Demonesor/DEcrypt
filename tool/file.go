package tool

import (
	"DEcrypt/config"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func Tempdir(c *config.Config) {
	//todo зробити обробку помилок

	c.Temp.BaseDir = os.TempDir()
	Log(c.Temp.BaseDir)

	// 2. Рахуємо хеш файлу
	c.Build.Hash, _ = getFileHash(c.Paths.AbsInput)
	Log(c.Build.Hash)

	// 3. Формуємо ім'я папки
	c.Temp.WorkDir = filepath.Base(c.Paths.AbsInput) + c.Build.Hash
	Log(c.Temp.WorkDir)

	// 4. Склеюємо ПОВНИЙ абсолютний шлях до нашої робочої папки всередині Temp

	c.Temp.WorkDir = filepath.Join(c.Temp.BaseDir, c.Temp.WorkDir)
	Log(c.Temp.WorkDir)

	// 5. Створюємо цю папку (MkdirAll не панікує, якщо папка вже існує)
	err := os.MkdirAll(c.Temp.WorkDir, 0755)
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
