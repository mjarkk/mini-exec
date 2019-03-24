package server

import (
	"crypto/sha1"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mjarkk/mini-exec/src/flags"
)

// sendErr returns true if there is an error
// and send the error to the connected client
func sendErr(c *gin.Context, err error, errorOverWrite ...string) bool {
	if err == nil {
		return false
	}
	fmt.Println("[ERROR]:", err.Error())
	errorToSend := err.Error()
	if len(errorOverWrite) > 0 {
		errorToSend = errorOverWrite[0]
	}
	c.JSON(400, gin.H{
		"status": false,
		"error":  errorToSend,
	})
	return true
}

// Start starts the webserver
func Start() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET("/", template)
	r.POST("/", handeler)

	r.Run(*flags.ServerLocation)
}

func template(c *gin.Context) {
	contentType := "text/html"

	data := []byte(Website)
	clientHash := strings.Replace(strings.Replace(c.GetHeader("If-None-Match"), "\"", "", -1), "W/", "", -1)
	fileHash := fmt.Sprintf("%x", sha1.Sum(data))
	c.Header("ETag", "\""+fileHash+"\"")
	resCode := 200

	if clientHash == fileHash {
		resCode = 304
	}

	c.Data(resCode, contentType, data)
}
