package entry

import (
	"Lv/internal/bootstrap"
	"Lv/internal/server"
)

//func usage(exitCode int, extraMessage ...interface{}) {
//	commands := console.GetCommands()
//
//	builder := new(strings.Builder)
//	builder.WriteString("⚙  可執行的指令 (command name)")
//	builder.WriteRune('\n')
//	commandName := "<none>"
//	for cmd := range commands {
//		command := commands[cmd]
//		commandName = cmd
//		builder.WriteString(fmt.Sprintf("		✏ %s %s\n", cmd, command.Description))
//	}
//
//	fmt.Printf(`
//	📖 環境說明 📖
//
//	需傳入以下環境變數：
//
//	⚙  APP_ENV : 專案環境
//		✏ docker 容器開發
//		✏ local 本機開發
//		✏ qatest 測試
//		✏ prod 正式
//
//	--------------
//
//	📖 指令說明 📖
//
//	⚙  主要指令
//		✏ server   運行伺服器
//		✏ schedule 運行背景排程
//		✏ run [command name] 執行指定命令
//
//	%s
//
//	📌  舉例： APP_ENV=local ./Lv server
//	📌  舉例： APP_ENV=local ./Lv schedule
//	📌  舉例： APP_ENV=local ./Lv run %s
//
//`, builder.String(), commandName)
//
//	if len(extraMessage) > 0 {
//		fmt.Println(extraMessage...)
//	}
//
//	os.Exit(exitCode)
//}

func Run() {

	//if bootstrap.GetAppEnv() == "" {
	//	usage(0)
	//} else {
	//	bootstrap.WriteLog("INFO", fmt.Sprintf("⚙  APP_ENV: %s", bootstrap.GetAppEnv()))
	//}
	//args := os.Args
	//if len(args) < 2 {
	//	usage(0)
	//	return
	//}

	// 設定優雅結束程序
	bootstrap.SetupGracefulSignal()
	server.Run()
}
