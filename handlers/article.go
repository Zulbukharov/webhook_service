package handlers

import (
	"fmt"
	"os"

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
	Old map[string]interface{} `json:"old"`
	New map[string]interface{} `json"new" binding:"required"`
}

type EventStruct struct {
	Operation string    `json:"op" binding:"required"`
	Data      EventData `json:"data" binding:"required"`
}

type HasuraEvent struct {
	Event *EventStruct `json:"event" binding:"required"`
	Table *TableStruct `json:"table" binding:"required"`
}

func Article(c *gin.Context) {
	header := c.GetHeader("Authorization")
	fmt.Println(header)
	if header != os.Getenv("AUTO") {
		c.JSON(400, gin.H{"error": "wrong header"})
		return
	}
	var body HasuraEvent
	err := c.BindJSON(&body)
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{"error": err})
		return
	}
	var message = "cannot process request"
	if body.Table.Name == "note" {
		message = fmt.Sprintf("New note %v inserted, with data: %v", body.Event.Data.New["id"], body.Event.Data.New["text"])
	}
	fmt.Println(message)
	c.JSON(200, gin.H{"message": message})
	return
}
