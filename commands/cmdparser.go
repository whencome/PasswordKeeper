package commands

import (
	"strings"
)

// 执行命令
func ExecCommand(command string, args []string) {
	cmd := strings.ToUpper(command)
	switch cmd {
	// 显示帮助
	case "HELP":
		showHelp()
	// 同步配置
	case "SYNC":
		syncConfigs()
	// 初始化
	case "INIT":
		initEnv()
	// 获取密码（复制到剪贴板，不直接展示）
	case "GET":
		getPassword(args)
	// 获取密码，并显示明文
	case "GETD":
		getPasswordDirect(args)
	// 设置密码
	case "SET":
		setPassword(args)
	// 展示所有项
	case "ITEMS":
		showItems()
	// 删除某个密码项
	case "DEL":
		deleteItem(args)
	// 测试项目
	case "TEST":
		execTest()
	}
}
