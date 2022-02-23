rm sync-site-bin
GOOS=linux GOARCH=amd64 go build -o sync-site-bin

rm sync-site-bin.zip
zip sync-site-bin.zip sync-site-bin

