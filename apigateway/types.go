package apigateway

import (
	"time"
)

type API struct {
	Id          string
	Name        string
	Version     string
	CreatedDate time.Time
	Description string
	//BinaryMediaTypes []string
}

type Resource struct {
	Id       string
	ParentId string
	Path     string
	PathPart string
	//ResourceMethods map[string]*Method
}
