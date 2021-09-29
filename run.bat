SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go mod tidy
go mod vendor
go build -o goFile main.go