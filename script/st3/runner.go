package st3

import (
	"DEcrypt/config"
	"DEcrypt/tool"
	"os"
	"path/filepath"
)

func runner(c *config.Config) {
	c.Build.Runner = `package main

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
	c.Build.RunnerPath = filepath.Join(c.Temp.WorkDir, "runner_main.go")
	err := os.WriteFile(c.Build.RunnerPath, []byte(c.Build.Runner), 0644)
	if err != nil {
		tool.Log("Помилка запису runner_main.go: " + err.Error())

	}
	return
}
