package tool

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
)

func getFileHash(input string) (string, error) {
	file, err := os.Open(input)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Створюємо хешер
	hasher := sha256.New()

	// Копіюємо вміст файлу прямо в хешер (це економно по пам'яті)
	if _, err := io.Copy(hasher, file); err != nil {
		return "", err
	}

	// Отримуємо хеш і переводимо в читабельний hex-рядок
	return hex.EncodeToString(hasher.Sum(nil)), nil
}
