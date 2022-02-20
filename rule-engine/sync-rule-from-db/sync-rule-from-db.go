package main

import (
	"encoding/json"
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
	log.Println("sync rule from db invoked")
	updateCache(retrieve())
}

func retrieve() (rules []protocol.Rule) {
	db, err := db.Connect()
	defer func() {
		_ = db.Close()
	}()
	if err != nil {
		log.Panic(err.Error())
	}

	q := fmt.Sprintf(`SELECT name, description, value, is_enabled FROM adminportal_rule;`)
	log.Println("debug: " + q)
	rows, err := db.Query(q)
	if err != nil {
		log.Panic(err.Error())
	}
	for rows.Next() {
		rule := protocol.Rule{}
		if err := rows.Scan(&rule.Name, &rule.Description, &rule.Value, &rule.IsEnabled); err != nil {
			log.Panic(err.Error())
		}
		bs, _ := json.Marshal(rule)
		log.Println("debug: " + string(bs))
		rules = append(rules, rule)
	}

	return rules
}

func updateCache(rules []protocol.Rule) {
	redisCli := cache.New(config.RuleCache)
	// redisCli.Del(config.RulePrefix + "*") // del all first
	for _, rule := range rules {
		value, _ := json.Marshal(rule)
		if err := redisCli.Set(fmt.Sprintf(config.RuleFormat, rule.Name), string(value), 0); err != nil {
			log.Panic(err.Error())
		}
	}
}
