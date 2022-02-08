rm ban-user-bin
GOOS=linux GOARCH=amd64 go build -o ban-user-bin

rm ban-user-bin.zip
zip ban-user-bin.zip ban-user-bin

