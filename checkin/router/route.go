package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda/messages"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	lambdaService "github.com/aws/aws-sdk-go/service/lambda"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"sg-stay-safe.org/checkin/config"
	"sg-stay-safe.org/checkin/protocol"
	"time"
)

func main() {
	log.Println("checkin in router service...")
	router := gin.Default()
	router.POST("/", route)

	router.Run(fmt.Sprintf(":%s", config.AWSForwardPort))
}

func route(c *gin.Context) {
	var event protocol.CheckInEvent
	if err := c.ShouldBindJSON(&event); err != nil {
		clientError(c, config.CodeMarshalError, err)
		return
	}

	code, err := checkin(event)
	if err != nil {
		log.Println(err.Error())
		serverError(c, code, err)
		return
	}
	success(c)
}

func checkin(event protocol.CheckInEvent) (int, error) {
	// use server timestamp
	event.Timestamp = time.Now().Unix()

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	client := lambdaService.New(sess, &aws.Config{Region: aws.String(config.AWSRegion)})

	payload, err := json.Marshal(event)
	if err != nil {
		log.Println(err.Error())
		return config.CodeMarshalError, err
	}

	lambdaChain := []string{
		config.SanitiseCheckinLambda,
		config.RecordCheckinLambda,
		config.ProduceCheckinMsgLambda,
	}
	for _, cmd := range lambdaChain {
		if resp, err := CallLambdaFunc(client, cmd, payload); err != nil {
			return resp.Code, err
		}
	}

	return config.CodeOK, nil
}

func CallLambdaFunc(client *lambdaService.Lambda, cmd string, payload []byte) (protocol.GeneralResponse, error) {
	var resp protocol.GeneralResponse

	result, err := client.Invoke(&lambdaService.InvokeInput{FunctionName: aws.String(cmd), Payload: payload})
	if err != nil {
		resp = protocol.GeneralResponse{
			Code: config.CodeInvokeLambdaError,
			Msg:  err.Error(),
		}
		return resp, err
	}

	log.Println("debug# payload: ", string(result.Payload))
	if result.FunctionError != nil {
		var invokeErr messages.InvokeResponse_Error
		_ = json.Unmarshal(result.Payload, &invokeErr)
		resp = protocol.GeneralResponse{
			Code: config.CodeInvokeLambdaError,
			Msg:  invokeErr.Message,
		}
		return resp, errors.New(resp.Msg)
	}

	err = json.Unmarshal(result.Payload, &resp)
	if err != nil {
		resp = protocol.GeneralResponse{
			Code: config.CodeUnmarshalError,
			Msg:  err.Error(),
		}
		log.Println(err.Error())
		return resp, err
	}
	log.Println(fmt.Sprintf("debug# resp msg: %s, code: %d", resp.Msg, resp.Code))

	if resp.Code != 0 {
		return resp, errors.New(resp.Msg)
	}

	return resp, nil
}

func serverError(c *gin.Context, errorCode int, err error) {
	resp := protocol.CheckInResponse{
		Code: errorCode,
		Msg:  err.Error(),
	}
	c.JSON(http.StatusInternalServerError, resp)
}

func clientError(c *gin.Context, errorCode int, err error) {
	resp := protocol.CheckInResponse{
		Code: errorCode,
		Msg:  err.Error(),
	}
	c.JSON(http.StatusBadRequest, resp)
}

func success(c *gin.Context) {
	resp := protocol.CheckInResponse{
		Code: config.CodeOK,
		Msg:  "check in successfully",
	}
	c.JSON(http.StatusOK, resp)
}
