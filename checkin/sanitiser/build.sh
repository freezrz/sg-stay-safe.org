rm sanitise-checkin-bin
GOOS=linux GOARCH=amd64 go build -o sanitise-checkin-bin

rm sanitise-checkin-bin.zip
zip sanitise-checkin-bin.zip sanitise-checkin-bin

