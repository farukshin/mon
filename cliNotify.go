package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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
	fmt.Println(cliNotifyListStr())
}

func cliNotifyListStr() string {
	resp, err := http.Get("http://localhost:1616/api/notify/list")
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return ""
		}
		var data []notification
		err = json.Unmarshal(bodyBytes, &data)
		if err != nil {
			return ""
		}
		var sb strings.Builder
		for _, s := range data {
			sb.WriteString(fmt.Sprintf("%s %s %s\n", s.UID, s.Type, s.Name))
		}
		return sb.String()
	}
	return ""
}

func cliNotifyAdd() {
	j := argsToJSON(os.Args[3:])
	data := []byte(j)
	r := bytes.NewReader(data)
	resp, err := http.Post("http://localhost:1616/api/notify/add", "application/json", r)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		fmt.Println(string(bodyBytes))
	}

}
func cliNotifyDelete() {
	j := argsToJSON(os.Args[3:])
	data := []byte(j)
	r := bytes.NewReader(data)
	resp, err := http.Post("http://localhost:1616/api/notify/delete", "application/json", r)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		fmt.Println(string(bodyBytes))
	}

}
func cliNotifyEdit() {
	j := argsToJSON(os.Args[3:])
	data := []byte(j)
	r := bytes.NewReader(data)
	resp, err := http.Post("http://localhost:1616/api/notify/edit", "application/json", r)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		fmt.Println(string(bodyBytes))
	}
}
