package main

import (
	ut "GoScanGun/utils"
	"fmt"
	"os"
)

const (
	MainPage = iota
	StudPage
	WorkPage
	ScanPage
)

var CurrentPage int

func main() {
	CurrentPage = MainPage
	for {
		ut.ShowMenu()
		fmt.Println("请输入你要操作的序号,并回车...")

		var inputCode int
		fmt.Scanf("%d", &inputCode)
		switch inputCode {
		case 1:
			fmt.Printf("输入是 1 新建作业收取任务%v\n", inputCode)
		case 2:
			fmt.Printf("输入是 2 列出已有任务%v\n", inputCode)
		case 3:
			fmt.Scanf("你要启动的作业收取任务是：&d4", &inputCode)
			// 启动 input 的任务编号
		case 4:
			os.Exit(0)
		}
	}

}
