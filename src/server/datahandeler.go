package server

import (
	"github.com/gin-gonic/gin"
)

func handeler(c *gin.Context) {
	var fullPostData PostData
	err := c.BindJSON(&fullPostData)
	if sendErr(c, err) {
		return
	}

	switch fullPostData.What {
	case "init":
		// Try to login

	case "update":
		// User pressed the update button

	case "check":
		// Check for new data

	case "reqToken":
		// Reqtoken generates a validation key and logs it in stdout
		err = GenValidation()
		if sendErr(c, err, "Failed to generate validation key") {
			return
		}

		c.JSON(200, gin.H{
			"status": true,
		})

	case "genToken":
		// logs the app key to stdout if the validation key matches
		var data GenToken
		err = DataTo(fullPostData, &data)
		if sendErr(c, err) {
			return
		}

		err = GenerateKey(data.ValidationKey)
		if sendErr(c, err) {
			return
		}

		c.JSON(200, gin.H{
			"status": true,
		})

	default:
		c.JSON(404, gin.H{
			"status": false,
			"err":    "Not a valid / not found: \"What\"",
		})
	}
}
