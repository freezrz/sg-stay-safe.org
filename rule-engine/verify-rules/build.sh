rm verify-rules-bin
GOOS=linux GOARCH=amd64 go build -o verify-rules-bin

rm verify-rules-bin.zip
zip verify-rules-bin.zip verify-rules-bin

