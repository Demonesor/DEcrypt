package script

import (
	"DEcrypt/config"
	"DEcrypt/tool"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func Pac(c *config.Config) {

	c.Exepath, c.Err = os.Executable()
	if c.Err != nil {
		fmt.Printf("Не вдалося отримати шлях бінарника: %v\n", c.Err)
		return
	}
	c.Exedir = filepath.Dir(c.Exepath) // Це папка, де лежить наш DEcrypt.exe

	// 2. Створюємо папку для завантажень (кешу) ПОРУЧ з нашим бінарником
	c.Datadir = filepath.Join(c.Exedir, "pdata")
	c.Err = os.MkdirAll(c.Datadir, 0755) // Створить папочку pdata, якщо її немає

	// 3. Отримуємо шлях до тимчасової робочої папки (вона в системному Temp)
	tool.Tempdir(c)

	// Копіюємо твій test.py в тимчасову папку
	targetPyFile := filepath.Join(c.Temp, filepath.Base(c.Input))
	tool.CopyFile(c.Input, targetPyFile)

	// 4. Завантажуємо python.zip в pdata (якщо його там немає)
	tool.Log("Downloading python...")
	zipPath := pythonD(c.Datadir) // Передаємо правильну папку кешу

	// Розпаковуємо Python ТИМЧАСОВО в робочу папку для тестів
	tempPythonDir := filepath.Join(c.Temp, "python")
	tool.Unzip(zipPath, tempPythonDir)

	// 5. Завантажуємо go.zip в pdata (якщо його там немає)
	tool.Log("Downloading go...")
	goPath := goD(c.Datadir)

	// Розпаковуємо TinyGo тимчасово в робочу папку
	tempTinygoDir := filepath.Join(c.Temp)
	tool.Unzip(goPath, tempTinygoDir)

	// 6. Запуск тестів
	tool.Log("Running tests...")

	// Шлях до тимчасового python.exe, який ми щойно розпакували в Temp
	interpreter := filepath.Join(tempPythonDir, "python.exe")

	// Запускаємо тест-ран
	TestRun(targetPyFile, interpreter)

	// ... Твій успішний TestRun пройшов тут ...

	tool.Log("Тести успішні! Починаю генерацію фінального бінарника...")

	runnerPath := runner(c.Temp)
	// 3. Копіюємо python.zip з pdata прямо в c.Temp, щоб embed його побачив поруч із runner_main.go

	c.Err = tool.CopyFile(filepath.Join(c.Datadir, "python.zip"), filepath.Join(c.Temp, "python.zip"))
	if c.Err != nil {
		tool.Log("Помилка копіювання python.zip для embed: ", c.Err)
		return
	}

	// Копіюємо скрипт як "script.py" в папку збірки
	c.Err = tool.CopyFile(c.Input, filepath.Join(c.Temp, "script.py"))

	// 4. КОМПІЛЯЦІЯ ЧЕРЕЗ GO
	tool.Log("Запускаю Go компіляцію...")

	// Визначаємо шлях, куди зберегти готовий скомпільований файл (наприклад, на робочий стіл)
	c.FinalExename = filepath.Base(c.Input) + "_packed.exe"
	c.FinalExePath = filepath.Join(c.Output, c.FinalExename)

	// Викликаємо встановлений у temp розпакований Go
	goExe := filepath.Join(c.Temp, "go", "bin", "go.exe")

	// Команда збірки: go build -ldflags="-s -w" -o <вихідний_exe> <шлях_до_runner_main.go>
	cmdBuild := exec.Command(goExe, "build", "-ldflags=-s -w", "-o", c.FinalExePath, runnerPath)
	cmdBuild.Dir = c.Temp // Збираємо прямо в temp папці, де лежить наш runner_main.go та вшиті файли

	cmdBuild.Stdout = os.Stdout
	cmdBuild.Stderr = os.Stderr

	tool.Log(c.Log, c.Output, c.TempDir, c.Temp, c.Input, c.FinalExePath, c.FinalExename, c.Datadir)

	c.Err = cmdBuild.Run()
	if c.Err != nil {
		tool.Log("Помилка компіляції Go: ", c.Err)
		return
	}

	tool.Log("==================================================")
	tool.Log("🎉 БІЛД ЗАВЕРШЕНО УСПІШНО! Файл створено:")
	tool.Log(c.FinalExePath)
	tool.Log("==================================================")
	return
}

func Clear(c *config.Config) {
	if c.Temp == "" || len(c.Temp) < 5 { // захист від коротких або порожніх шляхів
		return
	}
	os.RemoveAll(c.Temp)
}
