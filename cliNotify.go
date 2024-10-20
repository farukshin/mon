package main

import (
	"fmt"
	"os"
	"strings"
)

func cliNotify() {
	if os.Args[2] == "list" {
		cliNotifyList()
	} else if os.Args[2] == "add" {
		cliNotifyAdd()
	} else if os.Args[2] == "delete" {
		cliNotifyDelete()
	} else if os.Args[2] == "edit" {
		cliNotifyEdit()
	} else {
		notifyHome()
	}
}

func notifyHome() {
	fmt.Println(notifyHomeStr())
}

func notifyHomeStr() string {
	var sb strings.Builder
	sb.WriteString("Приложение: mon\n")
	sb.WriteString("\tСистема мониторинга\n\n")

	sb.WriteString("Строка запуска: mon notify [Опции]\n\n")

	sb.WriteString("Опции:\n")
	sb.WriteString("list - список всех уведомлений\n")
	sb.WriteString("add - добавить уведомление\n")
	sb.WriteString("delete - удалить уведомление\n")
	sb.WriteString("edit - изменить уведомление\n")

	sb.WriteString("Пример запуска:\n")
	sb.WriteString("./mon notify list\n")
	return sb.String()
}

func cliNotifyList() {

}
func cliNotifyAdd() {

}
func cliNotifyDelete() {

}
func cliNotifyEdit() {

}
