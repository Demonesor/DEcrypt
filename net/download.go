package net

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func DownloadFile(url string, filepath string) error {
	// 1. Робимо GET-запит
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Перевіряємо, чи сервер віддав 200 OK
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("помилка завантаження, статус: %s", resp.Status)
	}

	// 2. Створюємо файл
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// 3. Копіюємо вміст тіла відповіді у файл
	_, err = io.Copy(out, resp.Body)
	return err
}
