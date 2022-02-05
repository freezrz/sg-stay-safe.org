rm sanitise-checkin
GOOS=linux GOARCH=amd64 go build -o sanitise-checkin

rm sanitise-checkin.zip
zip sanitise-checkin.zip sanitise-checkin

