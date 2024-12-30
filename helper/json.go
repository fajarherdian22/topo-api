package helper

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func IsError(err error) {
	if err != nil {
		panic(err)
	}
}

// Encode Writer to JSON
func HandleEncodeWriteJson(c *gin.Context, WebResponse any) {
	c.Writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(c.Writer)
	err := encoder.Encode(WebResponse)
	IsError(err)
}

// Decode Request Body JSON
func HandleDecodeReqJson(c *gin.Context, dataStruct any) {
	decoder := json.NewDecoder(c.Request.Body)
	err := decoder.Decode(dataStruct)
	IsError(err)
	defer c.Request.Body.Close()
}
