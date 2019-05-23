package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

/*
                      +--------------+
              +-------+ HASURA EVENT +-----------+
              |       +--------------+           |
              v                                  v
      +-------+------+                   +-------+------+
      | TABLE STRUCT |                   | EVENT STRUCT |
      +-------+------+                   ++-------------+
              |                           |             |
              |                           v             v
+------+      |               +-----------++           ++----------+
| NAME +<-----+               | EVENT DATA |           | OPERATION |
+------+                      +------------+           +-----------+
                              |            |
                              v            v
                          +---+-+         ++----+
                          | OLD |         | NEW |
                          +-----+         +-----+

*/

type TableStruct struct {
	Name string `json:"name" binding:"required"`
}

type EventData struct {
	Old map[string]interface{} `json:"old" binding:"required"`
	New map[string]interface{} `json"new" binding:"required"`
}

type EventStruct struct {
	Operation string    `json:"op" binding:"required"`
	Data      EventData `json:"data" binding:"required"`
}

type HasuraEvent struct {
	Table *TableStruct `json:"table" binding:"required"`
	Event *EventStruct `json:"event" binding:"required"`
	Op    string       `json:"op" binding:"required"`
}

func Article(c *gin.Context) {
	body := &HasuraEvent{
		Table: &TableStruct{},
		Event: &EventStruct{},
	}
	err := c.BindJSON(&body)
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{"error": err})
		return
	}
	if err != nil {
		message := map[string]string{
			"message": "Unable to parse Hasura Event",
		}
		c.JSON(400, message)
		return
	}
	var message = "cannot process request"
	if body.Table.Name == "article" {
		message = fmt.Sprintf("New note %v inserted, with data: %v", body.Event.Data.New["id"], body.Event.Data.New["title"])
	}
	c.JSON(200, gin.H{"message": message})
	return
}
