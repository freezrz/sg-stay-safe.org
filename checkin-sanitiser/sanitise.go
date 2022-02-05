package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"net/http"
	"sg-stay-safe.org/checkin-sanitiser/protocol"
)

func main() {
	lambda.Start(Handler)
}

func Handler(ctx context.Context, event protocol.CheckInEvent) (protocol.SanitiserResponse, error) {
	fmt.Println("checkin event sanitise handler invoked")
	if event.AnonymousId == "" || event.PlaceId == "" {
		return protocol.SanitiserResponse{Code: http.StatusNotAcceptable, Msg: "not accepted"}, errors.New("not accepted")
	}
	return protocol.SanitiserResponse{Code: http.StatusOK, Msg: "sanitise ok..."}, nil
}
