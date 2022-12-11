build:
	go build -ldflags "-s -w"

checksum:
	sha256sum -b gatewayd-plugin-test
