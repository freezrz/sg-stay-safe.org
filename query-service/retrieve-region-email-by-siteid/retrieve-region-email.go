package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
	"sg-stay-safe.org/pkg/db"
	"sg-stay-safe.org/protocol"
)

func main() {
	lambda.Start(Handler)
}

func Handler(ctx context.Context, email protocol.ViolationEvent) (string, error) {
	log.Println("retrieve region email invoked")
	return retrieve(email.CheckInEvent.SiteId), nil
}

func retrieve(siteId string) (recipient string) {
	db, err := db.Connect()
	defer func() {
		_ = db.Close()
	}()
	if err != nil {
		log.Panic(err.Error())
	}

	q := fmt.Sprintf(`SELECT adminportal_region.email FROM adminportal_region, adminportal_site WHERE adminportal_site.site_id='%s' AND adminportal_site.region_id=adminportal_region.id;`, siteId)
	log.Println("debug: " + q)
	row := db.QueryRow(q)
	if err := row.Scan(&recipient); err != nil {
		recipient = "sg.stay.safe.org@gmail.com" // by default will report this to admin
	}
	log.Println("debug: " + recipient)

	return recipient
}
