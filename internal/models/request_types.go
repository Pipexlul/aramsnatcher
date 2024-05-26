package models

type RequestData map[string]interface{}

type RequestType string

const (
	GET    RequestType = "GET"
	POST   RequestType = "POST"
	DELETE RequestType = "DELETE"
	PATCH  RequestType = "PATCH"
)
