build:
	go build -o unitip-provider *.go
	mv unitip-provider "$(HOME)/Library/Application Scripts/tanin.tip/provider.script"
