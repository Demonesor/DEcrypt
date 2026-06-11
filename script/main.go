package script

import (
	"DEcrypt/config"
	"DEcrypt/script/st1"
	"DEcrypt/script/st2"
	"DEcrypt/tool"
	"os"
	"os/exec"
	"path/filepath"
)

func Start(c *config.Config) {

	// STEP 1 - преднастройка шукаємо знаходженя бінарника і создаєм папку кеша для рантайму і скачуєм його
	tool.Log("start stage1")
	st1.Start(c)
	// STEP 2 - копіюємо бінарники в времену папку
	tool.Log("start stage2")
	st2.Start(c)

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
	c.PathGo = filepath.Join(c.Temp, "go", "bin", "go.exe")

	// Команда збірки: go build -ldflags="-s -w" -o <вихідний_exe> <шлях_до_runner_main.go>
	cmdBuild := exec.Command(c.PathGo, "build", "-ldflags=-s -w", "-o", c.FinalExePath, runnerPath)
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
