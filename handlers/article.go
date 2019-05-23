package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

type TableStruct struct {
	Name string `json:"name"`
}

type EventData struct {
	Old map[string]interface{} `json:"old"`
	New map[string]interface{} `json"new"`
}

type EventStruct struct {
	Operation string    `json:"op"`
	Data      EventData `json:"data"`
}

type HasuraEvent struct {
	Table *TableStruct `json:"table"`
	Event *EventStruct `json:"event"`
	Op    string       `json:"op"`
}

func Article(c *gin.Context) {
	body := &HasuraEvent{
		Table: &TableStruct{},
		Event: &EventStruct{},
	}
	x, _ := ioutil.ReadAll(c.Request.Body)
	err := json.Unmarshal([]byte(x), body)
	if err != nil {
		message := map[string]string{
			"message": "Unable to parse Hasura Event",
		}
		c.JSON(400, message)
		return
	}
	var message = "cannot process request"
	var data = body.Event.Data
	if body.Table.Name == "article" {
		message = fmt.Sprintf("New note %v inserted, with data: %v", data.New["id"], data.New["title"])
	}
	c.JSON(200, gin.H{"message": message})
	return
}
