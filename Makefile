NAME=skip-brief-syscalls

.PHONY: binary gofmt test run clean

binary:
	go build -ldflags \
		"-X main.version=`git describe --tags --always --dirty`" \
		-o bin/$(NAME) this_module/main

gofmt:
	go fmt ./...

test:
	go test -cover -v ./...

run: binary
	strace -tT -o '!./bin/skip-brief-syscalls -d 0.5s' sleep 1

clean:
	rm -r bin
