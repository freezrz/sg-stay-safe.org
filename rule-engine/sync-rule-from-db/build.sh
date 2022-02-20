rm sync-rule-from-db-bin
GOOS=linux GOARCH=amd64 go build -o sync-rule-from-db-bin

rm sync-rule-from-db-bin.zip
zip sync-rule-from-db-bin.zip sync-rule-from-db-bin

