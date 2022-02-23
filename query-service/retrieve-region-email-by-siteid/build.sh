rm retrieve-region-email-bin
GOOS=linux GOARCH=amd64 go build -o retrieve-region-email-bin

rm retrieve-region-email-bin.zip
zip retrieve-region-email-bin.zip retrieve-region-email-bin

