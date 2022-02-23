rm sync-user-bin
GOOS=linux GOARCH=amd64 go build -o sync-user-bin

rm sync-user-bin.zip
zip sync-user-bin.zip sync-user-bin

