build:
	GOOS=linux GOARCH=386 go build -o ../bin/legendsbrowser-linux-386
	GOOS=linux GOARCH=amd64 go build -o ../bin/legendsbrowser-linux-x64
	GOOS=windows GOARCH=386 go build -o ../bin/legendsbrowser-386.exe
	GOOS=windows GOARCH=amd64 go build -o ../bin/legendsbrowser-x64.exe
	GOOS=darwin GOARCH=arm64 go build -o ../bin/legendsbrowser-macos-m1

buildMacOs:
	mkdir -p /tmp/build/macos
	cp -r macos/LegendsBrowser.app /tmp/build/macos/
	GOOS=darwin GOARCH=amd64 go build -o /tmp/build/macos/LegendsBrowser.app/Contents/MacOS/LegendsBrowser
	genisoimage -D -V "LegendsBrowser" -no-pad -r -apple -o /tmp/build/uncompressed.dmg /tmp/build/macos
	dmg dmg /tmp/build/uncompressed.dmg ../bin/LegendsBrowser.dmg

runLast:
	go run main.go -c conf.json -l=true
