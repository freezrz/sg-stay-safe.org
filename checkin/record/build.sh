rm record-checkin
GOOS=linux GOARCH=amd64 go build -o record-checkin

rm record-checkin.zip
zip record-checkin.zip record-checkin

