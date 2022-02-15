rm route-checkin
GOOS=linux GOARCH=amd64 go build -o route-checkin

rm route-checkin.zip
zip route-checkin.zip route-checkin

