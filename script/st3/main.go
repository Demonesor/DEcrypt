package st3

import (
	"DEcrypt/config"
	"DEcrypt/tool"
	"path/filepath"
)

func Start(c *config.Config) {
	// копіюємо код ранера в файл в робочой дерективі
	runner(c)

	// 3. Копіюємо python.zip з data прямо в c.Temp, щоб embed його побачив поруч із runner_main.go

	Err := tool.CopyFile(filepath.Join(c.Paths.DataDir, "python.zip"), filepath.Join(c.Temp.WorkDir, "python.zip"))
	if Err != nil {
		tool.Log("Помилка копіювання python.zip для embed: ", Err)
		return
	}

	// Копіюємо скрипт як "script.py" в папку збірки
	Err = tool.CopyFile(c.Paths.AbsInput, filepath.Join(c.Temp.WorkDir, "script.py"))

	// Визначаємо шлях, куди зберегти готовий скомпільований файл (наприклад, на робочий стіл)
	c.Build.FinalExeName = filepath.Base(c.Paths.AbsInput) + "_packed.exe"
	c.Build.FinalExePath = filepath.Join(c.Paths.Output, c.Build.FinalExeName)
}
