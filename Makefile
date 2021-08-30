.PHONY: all build run gotool clean help


BINARY="myexample"

all: gotool build


build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -m" -o ./bin/${BINARY}


run:
	@go run ./main.go

gotool:
	go fmt ./
	go vet ./



clean:
	@if [ -f ${BINARY} ]; then rm ${BINARY} ; fi


help:
	@echo "make - 格式化GO代码，并且编译生成二进制文件"
	@echo "make build - 编译GO代码，生成二进制文件"
	@echo "make run - 直接运行GO代码"
	@echo "make clean - 删除二进制文件"
	@echo "make gotool - 运行GO工具 'fmt' 和 'vet'"

