buildLinux:
	GOOS=linux GOARCH=amd64 go build -o ../bin/legendsbrowser-linux-x64

buildWindows:
	GOOS=windows GOARCH=amd64 go build -o ../bin/legendsbrowser-win-x64.exe

buildMacOs:
	GOOS=darwin GOARCH=amd64 go build -o ../bin/legendsbrowser-mac
	GOOS=darwin GOARCH=arm64 go build -o ../bin/legendsbrowser-mac-m1

buildMacOsApp:
	mkdir -p /tmp/build/macos
	cp -r macos/LegendsBrowser.app /tmp/build/macos/
	GOOS=darwin GOARCH=amd64 go build -o /tmp/build/macos/LegendsBrowser.app/Contents/MacOS/LegendsBrowser
	genisoimage -D -V "LegendsBrowser" -no-pad -r -apple -o /tmp/build/uncompressed.dmg /tmp/build/macos
	dmg dmg /tmp/build/uncompressed.dmg ../bin/LegendsBrowser.dmg

build: buildLinux buildWindows buildMacOs

runLast:
	go run main.go -c conf.json -l=true
