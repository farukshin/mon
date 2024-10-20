package main

import (
	"errors"
	"net/http"
	"os"
)

type notification struct {
	UID    string            `json:"uid"`
	Type   string            `json:"type"`
	Name   string            `json:"name"`
	Params map[string]string `json:"params"`
}

func (ntf *notification) send(mes string) error {

	if ntf.Type == "telegram" {
		MON_TELEGRAM_BOT_TOKEN, err := ntf.ntfParamOrEnv("MON_TELEGRAM_BOT_TOKEN", "MON_TELEGRAM_BOT_TOKEN")
		if err != nil {
			return err
		}
		MON_TELEGRAM_CHAT_ID, err := ntf.ntfParamOrEnv("MON_TELEGRAM_CHAT_ID", "MON_TELEGRAM_CHAT_ID")
		if err != nil {
			return err
		}
		ntf.sendMessageTG(mes, MON_TELEGRAM_BOT_TOKEN, MON_TELEGRAM_CHAT_ID)
	}
	return errors.New("Неизвестный тип уведомления")
}

func (ntf *notification) ntfParamOrEnv(param string, envName string) (string, error) {
	res, ok := ntf.Params[param]
	if !ok {
		envVal := os.Getenv(envName)
		if envVal == "" {
			return "", errors.New("Не указан параметр " + param)
		}
		return envVal, nil
	}
	return res, nil
}

func (app *application) addNotify(n *notification) {
	uid, _ := genUID()
	n.UID = uid
	app.Notifications = append(app.Notifications, *n)
}

func (app *application) deleteNotifyByID(uid string) error {

	for i, s := range app.Notifications {
		if s.UID == uid {
			copy(app.Notifications[i:], app.Notifications[i+1:])
			app.Notifications = app.Notifications[:len(app.Notifications)-1]
		}
	}
	return nil
}

func (app *application) getNotifyByUID(uid string) *notification {
	for i, _ := range app.Notifications {
		if app.Notifications[i].UID == uid {
			return &app.Notifications[i]
		}
	}
	return nil
}

func (ntf *notification) sendMessageTG(msg string, botToken string, chatID string) error {

	url := "https://api.telegram.org/bot" + botToken + "/sendMessage?chat_id=" + chatID + "&disable_web_page_preview=1&text=" + msg
	resp, err := http.Get(url)
	defer resp.Body.Close()
	return err
}
