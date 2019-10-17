INSTALL_PATH=/opt/cfupdater
BIN=cfupdater_linux_amd64
GOBUILD=CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o bin/$(BIN) cmd/updater.go

build:
	$(GOBUILD)