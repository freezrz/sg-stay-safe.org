package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
	"sg-stay-safe.org/config"
	"sg-stay-safe.org/pkg/cache"
	"sg-stay-safe.org/pkg/db"
	"sg-stay-safe.org/protocol"
)

func main() {
	lambda.Start(Handler)
}

func Handler() {
	log.Println("sync site from db invoked")
	updateCache(retrieve())
}

func retrieve() (sites []protocol.Site) {
	db, err := db.Connect()
	defer func() {
		_ = db.Close()
	}()
	if err != nil {
		log.Panic(err.Error())
	}

	q := fmt.Sprintf(`SELECT site_id FROM adminportal_site WHERE should_ban=1;`) // TODO: optimise the query with pagination.
	log.Println("debug: " + q)
	rows, err := db.Query(q)
	if err != nil {
		log.Panic(err.Error())
	}
	for rows.Next() {
		site := protocol.Site{}
		if err := rows.Scan(&site.SiteId); err != nil {
			log.Panic(err.Error())
		}
		log.Println("debug: " + site.SiteId)
		sites = append(sites, site)
	}

	return sites
}

func updateCache(sites []protocol.Site) {
	redisCli := cache.New(config.BanCache)
	for _, site := range sites {
		if err := redisCli.Set(fmt.Sprintf(config.BanSiteFormat, site.SiteId), "1", 24*60+60); err != nil { // let it expire after 24+1hours so no need to do daily cleanup
			log.Println(err.Error()) // TODO: enhance it with alert metrics
		}
	}
}
