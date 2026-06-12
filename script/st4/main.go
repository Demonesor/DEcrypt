package st4

import (
	"DEcrypt/config"
	"DEcrypt/tool"
	"os"
	"os/exec"
	"path/filepath"
)

func Start(c *config.Config) {
	// Викликаємо встановлений у temp розпакований Go

	c.Build.GOexe = filepath.Join(c.Paths.DataDir, "go", "bin", "go.exe")

	// Команда збірки: go build -ldflags="-s -w" -o <вихідний_exe> <шлях_до_runner_main.go>
	cmdBuild := exec.Command(c.Build.GOexe, "build", "-ldflags=-s -w", "-o", c.Build.FinalExePath, c.Build.RunnerPath)
	cmdBuild.Dir = c.Temp.WorkDir

	cmdBuild.Stdout = os.Stdout
	cmdBuild.Stderr = os.Stderr

	err := cmdBuild.Run()
	if err != nil {
		tool.Log("Помилка компіляції Go: ", err)
		return
	}

	tool.Log("==================================================")
	tool.Log("🎉 БІЛД ЗАВЕРШЕНО УСПІШНО! Файл створено:")
	tool.Log(c.Build.FinalExePath)
	tool.Log("==================================================")

}
