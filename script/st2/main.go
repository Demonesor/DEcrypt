package st2

import (
	"DEcrypt/config"
	"DEcrypt/tool"
	"os"
	"path/filepath"
)

func Start(c *config.Config) {
	tool.Log("start stage2 in st2")
	//TODO зробити обробку ошибок всіх!!

	// Копіюємо твій test.py в тимчасову папку
	c.Temp.Mainfiletemp = filepath.Join(c.Temp.WorkDir, filepath.Base(c.Paths.AbsInput))
	tool.CopyFile(c.Temp.Mainfiletemp, c.Paths.AbsInput)
	tool.Log(c.Temp.Mainfiletemp)

	// Розпаковуємо Python ТИМЧАСОВО в робочу папку для тестів

	c.Deps.PyTempDir = filepath.Join(c.Temp.WorkDir, "python")
	tool.Unzip(c.Deps.PyZip, c.Deps.PyTempDir)
	tool.Log(c.Deps.PyTempDir)

	// Розпаковуємо Go тимчасово в робочу папку
	g := filepath.Join(c.Paths.DataDir, "go", "bin", "go.exe")
	if _, err := os.Stat(g); err == nil {
		println("found go.exe cont")
	} else {

		c.Deps.GoTempDir = filepath.Join(c.Paths.DataDir)
		tool.Unzip(c.Deps.GoZip, c.Deps.GoTempDir)
		tool.Log(c.Deps.GoTempDir)
	}

	// 6. Запуск тестів
	tool.Log("Running tests...")

	// Шлях до тимчасового python.exe, який ми щойно розпакували в Temp
	c.Deps.Interpreter = filepath.Join(c.Deps.PyTempDir, "python.exe")
	tool.Log(c.Deps.Interpreter)

	// Запускаємо тест-ран
	TestRun(c.Temp.Mainfiletemp, c.Deps.Interpreter)

	// ... Твій успішний TestRun пройшов тут ...

	tool.Log("Тести успішні! Починаю генерацію фінального бінарника...")

	tool.Log("end stage2 in st2 ")
}
