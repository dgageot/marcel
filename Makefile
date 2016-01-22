build: marcel.go
	GOOS=linux GOARCH=amd64 go build -ldflags "-extldflags -static" marcel.go
	docker build -t dgageot/marcel .
