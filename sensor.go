package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type sensor struct {
	UID           string         `json:"uid"`
	Name          string         `json:"name"`
	Kind          string         `json:"kind"`
	Target        string         `json:"target"`
	Time          int            `json:"time"`
	Life          int            `json:"life"`
	Expect        sensorResult   `json:"expect"`
	History       []history      `json:"history"`
	Notifications []notification `json:"notifications"`
}

type sensorResultMessage struct {
	sensorUID string
	result    history
}

type sensorResult struct {
	resInt int
}

type history struct {
	time   time.Time
	result sensorResult
}

func (app *application) getSensorByUID(uid string) *sensor {
	for _, s := range app.Sensors {
		if s.UID == uid {
			return &s
		}
	}
	return nil
}

func (sen *sensor) check(srm *sensorResultMessage) bool {
	return sen.Expect.resInt == srm.result.result.resInt
}

func (sr sensorResult) String() string {
	return fmt.Sprintf("%d", sr.resInt)
}

func checkJob(sen *sensor) sensorResultMessage {
	resp, _ := http.Get(sen.Target)
	if resp == nil {
		return sensorResultMessage{
			sensorUID: sen.UID,
			result:    history{},
		}
	}
	defer resp.Body.Close()
	statuscode := resp.StatusCode
	return sensorResultMessage{
		sensorUID: sen.UID,
		result: history{
			time: time.Now(),
			result: sensorResult{
				resInt: statuscode,
			},
		},
	}
}

func (sen *sensor) calcNewTic() {
	//todo
}

func (app *application) addSensor(s *sensor) {
	id := uuid.New()
	s.UID = id.String()
	app.Sensors = append(app.Sensors, *s)
}

func (app *application) deleteSensorByID(uid string) error {

	for i, s := range app.Sensors {
		if s.UID == uid {
			copy(app.Sensors[i:], app.Sensors[i+1:])
			app.Sensors = app.Sensors[:len(app.Sensors)-1]
		}
	}
	return nil
}
