build:
	go mod tidy && go build -ldflags "-s -w"

checksum:
	sha256sum -b gatewayd-plugin-template

update-all:
	go get -u ./...
