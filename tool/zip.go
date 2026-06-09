package tool

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

// Unzip розпакує твій python.zip (src) у вказану папку (dest)
func Unzip(src string, dest string) error {
	// 1. Відкриваємо zip-архів
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	// 2. Біжимо по всіх файлах і папках всередині архіву
	for _, f := range r.File {
		println(f.Name)
		// Безпечно збираємо шлях (пам'ятаєш про кросплатформність?)
		fpath := filepath.Join(dest, f.Name)

		// Якщо це папка — створюємо її і йдемо далі
		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		// Якщо це файл — про всяк випадок створюємо для нього батьківську папку
		if err := os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return err
		}

		// 3. Створюємо цільовий файл на диску з тими ж правами доступу, що були в zip
		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}

		// 4. Відкриваємо файл всередині zip і копіюємо вміст
		rc, err := f.Open()
		if err != nil {
			outFile.Close()
			return err
		}

		_, err = io.Copy(outFile, rc)

		// Закриваємо дескриптори відразу в циклі, щоб не переповнити пам'ять
		outFile.Close()
		rc.Close()

		if err != nil {
			return err
		}
	}
	return nil
}
