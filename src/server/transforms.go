package server

// PostData the post data that is send to the server
type PostData struct {
	What          string `json:"what"`          // the type of request
	Key           string `json:"key"`           // the request token
	ValidationKey string `json:"validationKey"` // the validation key
}

// GenToken is the data type if PostsData.What is "GenToken"
type GenToken struct {
	ValidationKey string `json:"validationKey"` // the validation key
}

// DataToGenToken transforms PostData to GenToken
func DataToGenToken(d PostData) GenToken {
	var i interface{}
	i = d
	return i.(GenToken)
	// return GenToken{
	// 	ValidationKey: d.ValidationKey,
	// }
}
