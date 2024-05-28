SET CGO_ENABLED=0
SET GOOS=windows
SET GOARCH=amd64
go build -o ./build/fofac_win.exe main.go

SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build -o ./build/fofac_linux main.go

SET CGO_ENABLED=0
SET GOOS=darwin
SET GOARCH=amd64
go build -o ./build/fofac_mac main.go
