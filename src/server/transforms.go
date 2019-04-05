package server

import "encoding/json"

// PostData the post data that is send to the server
type PostData struct {
	What          string `json:"what"`          // the type of request
	Key           string `json:"key"`           // the request token
	ValidationKey string `json:"validationKey"` // the validation key
}

// GenToken is the data type if the request type is "GenToken"
type GenToken struct {
	ValidationKey string `json:"validationKey"` // the validation key
}

// Init is the data type if the request type is "Init"
type Init struct {
	Key string `json:"key"` // the request token
}

// DataTo transforms PostData to some other type in here
func DataTo(d PostData, bindTo interface{}) error {
	data, err := json.Marshal(d)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &bindTo)
	if err != nil {
		return err
	}

	return nil
}
