package st2

import (
	"DEcrypt/tool"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func TestRun(file, interpreter string) bool {
	tool.Log("Running " + file)
	tool.Log(interpreter)
	// Додаємо "-u" для моментального виводу
	cmd := exec.Command(interpreter, "-u", file)

	// НАЙВАЖЛИВІШЕ: Змушуємо Go виконувати команду НАЧЕБТО ми зайшли в папку з Python
	// Беремо папку, де лежить сам python.exe
	cmd.Dir = filepath.Dir(interpreter)

	// Направляємо вивід
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		// Якщо Go не зміг навіть рушити процес з місця, ми побачимо чому
		fmt.Printf("\n[Go System Error]: %v\n", err)
		return false
	}

	return true
}
