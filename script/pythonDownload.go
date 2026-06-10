package script

import (
	"DEcrypt/tool"
	"DEcrypt/tool/net"
	"os"
	"path/filepath"
)

func pythonD(dir string) string {
	// 1. Безпечно збираємо шлях під будь-яку ОС
	zipPath := filepath.Join(dir, "python.zip")

	// 2. Перевіряємо, чи файл РЕАЛЬНО існує на диску
	if _, err := os.Stat(zipPath); err == nil {
		tool.Log("Found local python.zip")
		return zipPath
	}

	// 3. Якщо файлу немає (одержали помилку від os.Stat), то качаємо
	tool.Log("python.zip не знайдено, завантажую з мережі...")
	err := net.DownloadFile("https://www.python.org/ftp/python/3.14.5/python-3.14.5-embed-amd64.zip", zipPath)
	if err != nil {
		tool.Log("Помилка завантаження файлу: " + err.Error())
		return ""
	}

	return zipPath
}
