package st1

import (
	"DEcrypt/config"
	"DEcrypt/tool"
	"DEcrypt/tool/net"
	"os"
	"path/filepath"
)

func goD(c *config.Config) {
	// 1. Безпечно збираємо шлях
	c.Deps.GoZip = filepath.Join(c.Paths.DataDir, "go.zip")

	// 2. Перевіряємо, чи файл РЕАЛЬНО існує на диску
	if _, err := os.Stat(c.Deps.GoZip); err == nil {
		tool.Log("Found local go.zip")
		return
	}

	// 3. Якщо файлу немає (одержали помилку від os.Stat), то качаємо
	tool.Log("go.zip не знайдено, завантажую з мережі...")
	err := net.DownloadFile("https://go.dev/dl/go1.26.4.windows-amd64.zip", c.Deps.GoZip)
	if err != nil {
		tool.Log("Помилка завантаження файлу: " + err.Error())
		return
	} //todo зробити обробку ошибки ітнернету

	return
}

func pythonD(c *config.Config) {
	// 1. Безпечно збираємо шлях під будь-яку ОС
	c.Deps.PyZip = filepath.Join(c.Paths.DataDir, "python.zip")

	// 2. Перевіряємо, чи файл РЕАЛЬНО існує на диску
	if _, err := os.Stat(c.Deps.PyZip); err == nil {
		tool.Log("Found local python.zip")
		return
	}

	// 3. Якщо файлу немає (одержали помилку від os.Stat), то качаємо
	tool.Log("python.zip не знайдено, завантажую з мережі...")
	err := net.DownloadFile("https://www.python.org/ftp/python/3.14.5/python-3.14.5-embed-amd64.zip", c.Deps.PyZip)
	if err != nil {
		tool.Log("Помилка завантаження файлу: " + err.Error())
		return
	} //todo зробити обробку ошибки ітнернету

	return
}
