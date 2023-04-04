build:
	go mod tidy && go build -ldflags "-s -w"

checksum:
	sha256sum -b plugin-template-go

update-all:
	go get -u ./...
