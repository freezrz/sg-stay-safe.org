package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
	"net/http"
	"sg-stay-safe.org/checkin-upsert/pkg"
	"sg-stay-safe.org/checkin-upsert/protocol"
)

func main() {
	lambda.Start(Handler)
}

func Handler(ctx context.Context, event protocol.CheckInEvent) (protocol.SanitiserResponse, error) {
	fmt.Println("upsert event sanitise handler invoked")
	if event.AnonymousId == "" || event.PlaceId == "" {
		return protocol.SanitiserResponse{Code: http.StatusNotAcceptable, Msg: "not accepted"}, errors.New("not accepted")
	}
	redisCli := pkg.New()
	err := redisCli.Set(event.PlaceId, "999", 500)
	if err != nil {
		log.Println(err.Error())
	}
	return protocol.SanitiserResponse{Code: http.StatusOK, Msg: "sanitise ok..."}, nil
}
