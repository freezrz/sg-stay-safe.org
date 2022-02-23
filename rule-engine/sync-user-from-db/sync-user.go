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
	log.Println("sync user from db invoked")
	updateCache(retrieve())
}

func retrieve() (users []protocol.User) {
	db, err := db.Connect()
	defer func() {
		_ = db.Close()
	}()
	if err != nil {
		log.Panic(err.Error())
	}

	q := fmt.Sprintf(`SELECT anonymous_id FROM adminportal_user WHERE should_ban=1;`) // TODO: optimise the query with pagination.
	log.Println("debug: " + q)
	rows, err := db.Query(q)
	if err != nil {
		log.Panic(err.Error())
	}
	for rows.Next() {
		user := protocol.User{}
		if err := rows.Scan(&user.AnonymousId); err != nil {
			log.Panic(err.Error())
		}
		log.Println("debug: " + user.AnonymousId)
		users = append(users, user)
	}

	return users
}

func updateCache(users []protocol.User) {
	redisCli := cache.New(config.BanCache)
	for _, user := range users {
		if err := redisCli.Set(fmt.Sprintf(config.BanUserFormat, user.AnonymousId), "1", 24*60+60); err != nil { // let it expire after 24+1hours so no need to do daily cleanup
			log.Println(err.Error()) // TODO: enhance it with alert metrics
		}
	}
}
