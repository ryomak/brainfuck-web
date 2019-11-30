package main

import (
	"context"

	"github.com/ryomak/brainfuck-web/functions/bf"
)

var maxCommand = 4000

var setting = bf.Commands{
	NEXT:  "",
	PREV:  "",
	INC:   "",
	DEC:   "",
	WRITE: "",
	OPEN:  "",
	CLOSE: "",
}

var state = bf.InitState(setting, maxCommand)

type Event struct {
	Code string `json:"name"`
}

type Response struct {
	Result string `json:"result"`
}

func execureHandler(ctx context.Context, event Event) (Response, error) {
	result := state.Start(event.Code)
	return Response{result}, nil
}

func main() {
	lambda.Start(execureHandler)
}
