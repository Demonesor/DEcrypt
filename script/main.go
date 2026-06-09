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

	// 1. Формуємо текст runner_main.go (запихуємо шаблон вище в рядок)
	runnerCode := `package main

import (
    "archive/zip"
    "embed"
    "fmt"
    "io"
    "os"
    "os/exec"
    "path/filepath"
)

//go:embed python.zip script.py
var embeddedFiles embed.FS

// Чиста функція розпаковки на стандартній бібліотеці Go, щоб файл був автономним
func unzipBytes(zipPath, destDir string) error {
    r, err := zip.OpenReader(zipPath)
    if err != nil {
        return err
    }
    defer r.Close()

    for _, f := range r.File {
        fpath := filepath.Join(destDir, f.Name)
        if f.FileInfo().IsDir() {
            os.MkdirAll(fpath, os.ModePerm)
            continue
        }

        if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
            return err
        }

        outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
        if err != nil {
            return err
        }

        rc, err := f.Open()
        if err != nil {
            outFile.Close()
            return err
        }

        _, err = io.Copy(outFile, rc)
        outFile.Close()
        rc.Close()
        if err != nil {
            return err
        }
    }
    return nil
}

func main() {
    tempDir, err := os.MkdirTemp("", "decore_*")
    if err != nil {
        fmt.Println("Error creating temp dir:", err)
        return
    }
    defer os.RemoveAll(tempDir)

    // Вивантажуємо python.zip
    pyZipBytes, _ := embeddedFiles.ReadFile("python.zip")
    pyZipPath := filepath.Join(tempDir, "python.zip")
    os.WriteFile(pyZipPath, pyZipBytes, 0644)

    // Вивантажуємо скрипт
    scriptBytes, _ := embeddedFiles.ReadFile("script.py")
    scriptPath := filepath.Join(tempDir, "runtime.py")
    os.WriteFile(scriptPath, scriptBytes, 0644)

    // Розпакувуємо вбудованим методом
    targetPyDir := filepath.Join(tempDir, "python")
    err = unzipBytes(pyZipPath, targetPyDir)
    if err != nil {
        fmt.Println("[Runner Unzip Error]:", err)
        return
    }

    // Запуск
    interpreter := filepath.Join(targetPyDir, "python.exe")
    cmd := exec.Command(interpreter, "-u", scriptPath)
    
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    err = cmd.Run()
    if err != nil {
        fmt.Println("[Runner Execution Error]:", err)
    }
}
`

	// 2. Записуємо runner_main.go в нашу тимчасову папку збірки
	runnerPath := filepath.Join(tempDir, "runner_main.go")
	err = os.WriteFile(runnerPath, []byte(runnerCode), 0644)
	if err != nil {
		tool.Log("Помилка запису runner_main.go: " + err.Error())
		return
	}

	// 3. Копіюємо python.zip з pdata прямо в tempDir, щоб embed його побачив поруч із runner_main.go
	// Важливо: embed вміє читати файли ТІЛЬКИ з тієї ж папки, де лежить .go файл, або нижче!
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
