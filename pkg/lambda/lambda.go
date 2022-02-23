package lambdaPkg

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda/messages"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/lambda"
	"log"
	"sg-stay-safe.org/config"
	"sg-stay-safe.org/protocol"
)

func CallLambdaFunc(client *lambda.Lambda, cmd string, payload []byte) (protocol.GeneralResponse, []byte, error) {
	var resp protocol.GeneralResponse

	result, err := client.Invoke(&lambda.InvokeInput{FunctionName: aws.String(cmd), Payload: payload})
	if err != nil {
		resp = protocol.GeneralResponse{
			Code: config.CodeInvokeLambdaError,
			Msg:  err.Error(),
		}
		return resp, nil, err
	}

	log.Println("debug# payload: ", string(result.Payload))
	if result.FunctionError != nil {
		var invokeErr messages.InvokeResponse_Error
		_ = json.Unmarshal(result.Payload, &invokeErr)
		resp = protocol.GeneralResponse{
			Code: config.CodeInvokeLambdaError,
			Msg:  invokeErr.Message,
		}
		return resp, nil, errors.New(resp.Msg)
	}

	err = json.Unmarshal(result.Payload, &resp)
	if err != nil {
		resp = protocol.GeneralResponse{
			Code: config.CodeUnmarshalError,
			Msg:  err.Error(),
		}
		log.Println(err.Error())
		return resp, result.Payload, err
	}
	log.Println(fmt.Sprintf("debug# resp msg: %s, code: %d", resp.Msg, resp.Code))

	if resp.Code != 0 {
		return resp, result.Payload, errors.New(resp.Msg)
	}

	return resp, result.Payload, nil
}
