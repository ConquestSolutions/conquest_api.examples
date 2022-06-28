package conquest_api

import (
	"encoding/json"
	"fmt"
)

// The response was not OK, not JSON and not a standard ErrorResponse
// This is returned as an error, which you could handle elsewhere
type UnknownResponse struct {
	Status      int
	ContentType string
	Content     string
}

func (e UnknownResponse) Error() string {
	if e.Content == "" {
		return fmt.Sprint(e.Status)
	}
	return fmt.Sprintf(e.Content)
}

type ErrorResponse struct {
	ErrorCode        string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

func (e ErrorResponse) Error() string {
	if e.ErrorDescription == "" {
		return e.ErrorCode
	}
	return e.ErrorCode + ": " + e.ErrorDescription
}

// Empty is just an empty object. Use it when you don't expect or need the return value
type Empty struct {
}

// RawMessage is valid json that you can use to read something that can be anything: {} or [] or "" or 1 or true or null
// It's handy if you want to unpack one message into many values or want to handle the message in some other special way.
type RawMessage struct {
	Json json.RawMessage
}

type RequestError interface {
	Code() int
	Error() string
}
