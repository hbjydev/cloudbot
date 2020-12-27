all: build

build:
	go build -o cloudbot cmd/cloudbot/main.go

run:
	CLOUDBOT_IRC_NICK="cloudbot__" CLOUDBOT_IRC_CHANNEL="#cloudbot-devel" go run -race cmd/cloudbot/main.go
