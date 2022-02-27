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

func Handler(ctx context.Context, email protocol.ViolationEvent) (protocol.Region, error) {
	log.Println("retrieve region email invoked")
	return retrieve(email.CheckInEvent.SiteId), nil
}

func retrieve(siteId string) (region protocol.Region) {
	db, err := db.Connect()
	defer func() {
		_ = db.Close()
	}()
	if err != nil {
		log.Panic(err.Error())
	}

	q := fmt.Sprintf(`SELECT adminportal_region.name, adminportal_region.description, adminportal_region.email FROM adminportal_region, adminportal_site WHERE adminportal_site.site_id='%s' AND adminportal_site.region_id=adminportal_region.id;`, siteId)
	log.Println("debug: " + q)
	row := db.QueryRow(q)
	if err := row.Scan(&region.Name, &region.Description, &region.Email); err != nil {
		region.Name = "system-default-region"
		region.Description = "missing region info, using default region"
		region.Email = "sg.stay.safe.org@gmail.com" // by default will report this to admin
	}
	log.Println("debug: " + region.Email)

	return region
}
