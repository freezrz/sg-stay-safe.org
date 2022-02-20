rm record-checkin-bin
GOOS=linux GOARCH=amd64 go build -o record-checkin-bin

rm record-checkin-bin.zip
zip record-checkin-bin.zip record-checkin-bin

