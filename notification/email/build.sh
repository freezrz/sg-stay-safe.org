rm send-email-bin
GOOS=linux GOARCH=amd64 go build -o send-email-bin

rm send-email-bin.zip
zip send-email-bin.zip send-email-bin

