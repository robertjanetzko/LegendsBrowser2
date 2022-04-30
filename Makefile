build:
	cd backend && GOOS=linux GOARCH=386 go build -o ../bin/legendsbrowser-linux-386
	cd backend && GOOS=linux GOARCH=amd64 go build -o ../bin/legendsbrowser-linux-x64
	cd backend && GOOS=windows GOARCH=386 go build -o ../bin/legendsbrowser-386.exe
	cd backend && GOOS=windows GOARCH=amd64 go build -o ../bin/legendsbrowser-x64.exe
	cd backend && GOOS=darwin GOARCH=amd64 go build -o ../bin/legendsbrowser-macos-x64
	cd backend && GOOS=darwin GOARCH=arm64 go build -o ../bin/legendsbrowser-macos-m1

run:
	cd backend && go run main.go