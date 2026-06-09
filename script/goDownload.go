package script

import (
	"DEcrypt/net"
	"DEcrypt/tool"
	"os"
	"path/filepath"
)

func goD(dir string) string {
	// 1. Безпечно збираємо шлях під будь-яку ОС
	zipPath := filepath.Join(dir, "go.zip")

	// 2. Перевіряємо, чи файл РЕАЛЬНО існує на диску
	if _, err := os.Stat(zipPath); err == nil {
		tool.Log("Found local go.zip")
		return zipPath
	}

	// 3. Якщо файлу немає (одержали помилку від os.Stat), то качаємо
	tool.Log("go.zip не знайдено, завантажую з мережі...")
	err := net.DownloadFile("https://go.dev/dl/go1.26.4.windows-amd64.zip", zipPath)
	if err != nil {
		tool.Log("Помилка завантаження файлу: " + err.Error())
		return ""
	}

	return zipPath
}
