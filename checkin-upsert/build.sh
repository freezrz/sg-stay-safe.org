rm upsert-checkin
GOOS=linux GOARCH=amd64 go build -o upsert-checkin

rm upsert-checkin.zip
zip upsert-checkin.zip upsert-checkin

