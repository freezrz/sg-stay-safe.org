rm anti-fraud-bin
GOOS=linux GOARCH=amd64 go build -o anti-fraud-bin

rm anti-fraud-bin.zip
zip anti-fraud-bin.zip anti-fraud-bin # dummy update

