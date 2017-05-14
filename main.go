package main

import (
	"fmt"

	"github.com/samalba/lambda-gateway/apigateway"
)

func main() {
	// List APIS
	apiGw := apigateway.NewAPIGateway()
	apis, err := apiGw.ListAPIs()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%#v\n", apis)
	}

	// Get Resources for the first API
	ress, err := apiGw.GetResources(apis[0].Id)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%#v\n", ress)
	}
}
