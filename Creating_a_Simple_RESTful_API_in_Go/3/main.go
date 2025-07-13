package main

import "fmt"

type Request struct {
	ID    float64 `json:id`
	Value string  `json:value`
}

type Response struct {
	Message string `json:message`
	Ok      bool   `json:ok`
}


func Handler (request Request) (Response , error){
	return Response{
		Message: fmt.Sprintf("process request ID %f",request.ID),
		Ok:true,
	}, nil
}