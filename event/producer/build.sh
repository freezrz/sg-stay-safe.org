rm produce-event-bin
GOOS=linux GOARCH=amd64 go build -o produce-event-bin

rm produce-event-bin.zip
zip produce-event-bin.zip produce-event-bin

