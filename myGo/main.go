package main

import (
	"Lv/internal/entry"
	"os"

	_ "Lv/docs"
)

func init(){
	os.Setenv("TZ","Asia/Taiper")
}

// @title API文件
// @version 1.0.0
// @description  OO系統
// @termsOfService http://swagger.io/terms/
// @license.name Marvin
func main()  {
	entry.Run()
}