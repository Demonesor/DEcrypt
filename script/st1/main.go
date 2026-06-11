package st1

import (
	"DEcrypt/config"
	"DEcrypt/tool"
	"os"
	"path/filepath"
)

func Start(c *config.Config) {
	//todo зробити логуваня кроків

	tool.Log("start stage1 in st1")

	//пошук exe файла

	tool.Log("searching for exe file")
	c.Paths.ExePath, _ = os.Executable() //TODO зробити обробку ошибок
	tool.Log(c.Paths.ExePath)

	tool.Log("give me dir of exe")
	c.Paths.ExeDir = filepath.Dir(c.Paths.ExePath) //получаєм деректорію exe
	tool.Log(c.Paths.ExeDir)

	tool.Log("create pdata")
	c.Paths.DataDir = filepath.Join(c.Paths.ExeDir, "pdata") // робим папку даних около exe pdata
	_ = os.Mkdir(c.Paths.DataDir, 0755)                      //TODO зробити обробку ошибок
	tool.Log(c.Paths.DataDir)

	tool.Tempdir(c)

	tool.Log("Downloading python...")
	pythonD(c)
	tool.Log("Dowloading go...")
	goD(c)

	tool.Log("end stage1 in st1")

}
