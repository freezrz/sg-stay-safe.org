package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
	"sg-stay-safe.org/config"
	"sg-stay-safe.org/protocol"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

const (
	Sender = config.SystemEmailSender

	Subject = "sg-stay-safe: safe distance violation detected!!!"

	HtmlBody = "<h1>The platform just detected that below site/user is likely violating the safe distance!!! </h1>" +
		"<p>Site Info: %s</p>" +
		"<p>Anonymous Id: %s</p>" +
		"<p>Reason: %s</p>"

	TextBody = "Site Info: %s. Anonymous Id: %s"

	CharSet = "UTF-8"
)

func main() {
	lambda.Start(handler)
}

func handler(ctx context.Context, event protocol.ViolationEmail) (protocol.GeneralResponse, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(config.AWSRegion)},
	)

	svc := ses.New(sess)

	if event.AnonymousId == "" {
		event.AnonymousId = "N/A"
	}

	var siteInfo []byte
	if siteInfo, err = json.Marshal(event.Site); err != nil {
		siteInfo = []byte{}
	}

	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			CcAddresses: []*string{},
			ToAddresses: []*string{
				aws.String(event.Recipient),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String(CharSet),
					Data:    aws.String(fmt.Sprintf(HtmlBody, string(siteInfo), event.AnonymousId, event.Reason)),
				},
				Text: &ses.Content{
					Charset: aws.String(CharSet),
					Data:    aws.String(TextBody),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String(CharSet),
				Data:    aws.String(Subject),
			},
		},
		Source: aws.String(Sender),
	}

	result, err := svc.SendEmail(input)

	if err != nil {
		return protocol.GeneralResponse{Code: config.CodeNotificationSendEmailError, Msg: fmt.Sprintf("email send error: %s", err.Error())}, nil
	}

	log.Println(fmt.Sprintf("email sent to address: %s, msg id: %s", event.Recipient, *result.MessageId))
	log.Println(result)
	return protocol.GeneralResponse{Code: config.CodeOK, Msg: fmt.Sprintf("email sent to %s", event.Recipient)}, nil
}
