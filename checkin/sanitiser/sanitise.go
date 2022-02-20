package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
	"sg-stay-safe.org/config"
	"sg-stay-safe.org/protocol"
)

func main() {
	lambda.Start(Handler)
}

func Handler(ctx context.Context, event protocol.CheckInEvent) (protocol.GeneralResponse, error) {
	log.Println("sanitiser invoked")
	if event.AnonymousId == "" || event.SiteId == "" {
		return protocol.GeneralResponse{Code: config.CodeSanitiseError, Msg: "anonymous id and site id can't be empty"}, nil
	}
	return protocol.GeneralResponse{Code: config.CodeOK, Msg: "sanitise ok..."}, nil
}
