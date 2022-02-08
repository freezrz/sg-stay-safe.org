rm consume-event-bin
GOOS=linux GOARCH=amd64 go build -o consume-event-bin

rm consume-event-bin.zip
zip consume-event-bin.zip consume-event-bin

