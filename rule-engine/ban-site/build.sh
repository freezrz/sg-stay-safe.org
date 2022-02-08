rm ban-site-bin
GOOS=linux GOARCH=amd64 go build -o ban-site-bin

rm ban-site-bin.zip
zip ban-site-bin.zip ban-site-bin

