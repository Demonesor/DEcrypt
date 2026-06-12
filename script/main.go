package script

import (
	"DEcrypt/config"
	"DEcrypt/script/st1"
	"DEcrypt/script/st2"
	"DEcrypt/script/st3"
	"DEcrypt/script/st4"
	"DEcrypt/tool"
	"os"
)

func Start(c *config.Config) {

	// STEP 1 - преднастройка шукаємо знаходженя бінарника і создаєм папку кеша для рантайму і скачуєм його
	tool.Log("start stage1")
	st1.Start(c)
	// STEP 2 - копіюємо бінарники в времену папку
	tool.Log("start stage2")
	st2.Start(c)

	tool.Log("start stage3")
	st3.Start(c)

	tool.Log("start stage4")
	st4.Start(c)

	return
}

func Clear(c *config.Config) {
	if c.Temp.WorkDir == "" || len(c.Temp.WorkDir) < 5 { // захист від коротких або порожніх шляхів
		return
	}
	os.RemoveAll(c.Temp.WorkDir)
}
