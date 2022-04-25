build:
	cd backend && GOOS=linux GOARCH=386 go build -o ../bin/legendsbrowser-linux-386 main.go
	cd backend && GOOS=linux GOARCH=amd64 go build -o ../bin/legendsbrowser-linux-x64 main.go
	cd backend && GOOS=windows GOARCH=386 go build -o ../bin/legendsbrowser-386.exe main.go
	cd backend && GOOS=windows GOARCH=amd64 go build -o ../bin/legendsbrowser-x64.exe main.go
	cd backend && GOOS=darwin GOARCH=amd64 go build -o ../bin/legendsbrowser-macos-x64 main.go
	cd backend && GOOS=darwin GOARCH=arm64 go build -o ../bin/legendsbrowser-macos-m1 main.go

run:
	cd backend && go run main.go