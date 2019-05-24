package tests

import (
	"bytes"
	"fmt"
	"os"

	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Zulbukharov/webhook_service/routes"

	"github.com/bmizerany/assert"

	"github.com/gin-gonic/gin"
)

var (
	header = os.Getenv("AUTO")
)

// type Payload struct {
// 	Event struct {
// 		Op   string `json:"op"`
// 		Data struct {
// 			Old interface{} `json:"old"`
// 			New struct {
// 				Text string `json:"text"`
// 				ID   int    `json:"id"`
// 			} `json:"new"`
// 		} `json:"data"`
// 	} `json:"event"`
// 	Table struct {
// 		Schema string `json:"schema"`
// 		Name   string `json:"name"`
// 	} `json:"table"`
// }

func SetupRouter() *gin.Engine {
	router := gin.Default()
	gin.SetMode(gin.TestMode)
	routes.InitializeRoutes(router)
	return router
}

func main() {
	r := SetupRouter()
	r.Run(":8001")
}

var wrongPayload = `
{
	"event": {
        "data": {
            "old": null,
            "new": {
                "text": "new-entry",
                "id": 1
			}
        }
    },
    "table": {
        "name": "note"
    }
}
`

var correctPayload = `
{
	"event": {
		"op": "INSERT",
        "data": {
            "old": null,
            "new": {
                "text": "new-entry",
                "id": 1
			}
        }
    },
    "table": {
        "schema": "public",
        "name": "note"
	}
}
`

/**
* TestSignup
* Test user registration
*
* Must return response code 200
 */
func TestCorrectPost(t *testing.T) {
	testRouter := SetupRouter()

	req, err := http.NewRequest("POST", "/article", bytes.NewBufferString(correctPayload))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", header)

	if err != nil {
		fmt.Println(err)
	}
	resp := httptest.NewRecorder()

	testRouter.ServeHTTP(resp, req)
	assert.Equal(t, resp.Code, 200)
}

/**
* TestSignup
* Test user registration
*
* Must return response code 200
 */
func TestWrongPostWithoudOp(t *testing.T) {
	testRouter := SetupRouter()

	req, err := http.NewRequest("POST", "/article", bytes.NewBufferString(wrongPayload))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", header)

	if err != nil {
		fmt.Println(err)
	}
	resp := httptest.NewRecorder()
	fmt.Println(resp.Body)
	testRouter.ServeHTTP(resp, req)
	assert.Equal(t, resp.Code, 400)
}

func TestWrongPostWithoudData(t *testing.T) {
	testRouter := SetupRouter()

	payload := []byte(`{
		"event":{
			"op":"INSERT",
			"data":{
				old: null,
				new: null,
			},
		},
		"table":{
			"name":"note",
		},
	}`)
	// err := json.Unmarshal(payload, &data)
	req, err := http.NewRequest("POST", "/article", bytes.NewBufferString(string(payload)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", header)

	if err != nil {
		fmt.Println(err)
	}
	resp := httptest.NewRecorder()
	fmt.Println(resp.Body)
	testRouter.ServeHTTP(resp, req)
	assert.Equal(t, resp.Code, 400)
}

func TestWrongPostEmpty(t *testing.T) {
	testRouter := SetupRouter()

	payload := []byte(`{
		"event":{
		},
		"table":{
			"name":"note",
		},
	}`)
	// err := json.Unmarshal(payload, &data)
	req, err := http.NewRequest("POST", "/article", bytes.NewBufferString(string(payload)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", header)

	if err != nil {
		fmt.Println(err)
	}
	resp := httptest.NewRecorder()
	fmt.Println(resp.Body)
	testRouter.ServeHTTP(resp, req)
	assert.Equal(t, resp.Code, 400)
}
