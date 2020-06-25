build:
	go build -o unitip-provider *.go
	cp unitip-provider "$(HOME)/Library/Application Scripts/tanin.tip/provider.script"
