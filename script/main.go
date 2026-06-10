package script

import (
	"DEcrypt/tool"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func Pac(input string, output string) {
	// 1. Отримуємо шлях до нашого запущеного .exe
	exePath, err := os.Executable()
	if err != nil {
		fmt.Printf("Не вдалося отримати шлях бінарника: %v\n", err)
		return
	}
	exeDir := filepath.Dir(exePath) // Це папка, де лежить наш DEcrypt.exe

	// 2. Створюємо папку для завантажень (кешу) ПОРУЧ з нашим бінарником
	pdataDir := filepath.Join(exeDir, "pdata")
	os.MkdirAll(pdataDir, 0755) // Створить папочку pdata, якщо її немає

	// 3. Отримуємо шлях до тимчасової робочої папки (вона в системному Temp)
	tempDir := tool.Tempdir(input, output)

	// Копіюємо твій test.py в тимчасову папку
	targetPyFile := filepath.Join(tempDir, filepath.Base(input))
	tool.CopyFile(input, targetPyFile)

	// 4. Завантажуємо python.zip в pdata (якщо його там немає)
	tool.Log("Downloading python...")
	zipPath := pythonD(pdataDir) // Передаємо правильну папку кешу

	// Розпаковуємо Python ТИМЧАСОВО в робочу папку для тестів
	tempPythonDir := filepath.Join(tempDir, "python")
	tool.Unzip(zipPath, tempPythonDir)

	// 5. Завантажуємо go.zip в pdata (якщо його там немає)
	tool.Log("Downloading go...")
	goPath := goD(pdataDir)

	// Розпаковуємо TinyGo тимчасово в робочу папку
	tempTinygoDir := filepath.Join(tempDir)
	tool.Unzip(goPath, tempTinygoDir)

	// 6. Запуск тестів
	tool.Log("Running tests...")

	// Шлях до тимчасового python.exe, який ми щойно розпакували в Temp
	interpreter := filepath.Join(tempPythonDir, "python.exe")

	// Запускаємо тест-ран
	TestRun(targetPyFile, interpreter)

	// ... Твій успішний TestRun пройшов тут ...

	tool.Log("Тести успішні! Починаю генерацію фінального бінарника...")

	runnerPath := runner(tempDir)
	// 3. Копіюємо python.zip з pdata прямо в tempDir, щоб embed його побачив поруч із runner_main.go

	err = tool.CopyFile(filepath.Join(pdataDir, "python.zip"), filepath.Join(tempDir, "python.zip"))
	if err != nil {
		tool.Log("Помилка копіювання python.zip для embed: " + err.Error())
		return
	}

	// Копіюємо скрипт як "script.py" в папку збірки
	err = tool.CopyFile(input, filepath.Join(tempDir, "script.py"))

	// 4. КОМПІЛЯЦІЯ ЧЕРЕЗ GO
	tool.Log("Запускаю Go компіляцію...")

	// Визначаємо шлях, куди зберегти готовий скомпільований файл (наприклад, на робочий стіл)
	finalExeName := filepath.Base(input) + "_packed.exe"
	finalExePath := filepath.Join(output, finalExeName)

	// Викликаємо встановлений у temp розпакований Go
	goExe := filepath.Join(tempDir, "go", "bin", "go.exe")

	// Команда збірки: go build -ldflags="-s -w" -o <вихідний_exe> <шлях_до_runner_main.go>
	cmdBuild := exec.Command(goExe, "build", "-ldflags=-s -w", "-o", finalExePath, runnerPath)
	cmdBuild.Dir = tempDir // Збираємо прямо в temp папці, де лежить наш runner_main.go та вшиті файли

	cmdBuild.Stdout = os.Stdout
	cmdBuild.Stderr = os.Stderr

	err = cmdBuild.Run()
	if err != nil {
		tool.Log("Помилка компіляції Go: " + err.Error())
		return
	}

	tool.Log("==================================================")
	tool.Log("🎉 БІЛД ЗАВЕРШЕНО УСПІШНО! Файл створено:")
	tool.Log(finalExePath)
	tool.Log("==================================================")
	return
}
func Clear(input string, output string) {
	temp := tool.Tempdir(input, output)
	os.RemoveAll(temp)
	return
}
