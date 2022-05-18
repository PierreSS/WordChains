all: go

exec = wordchains

go: main.go
	go test ./
	go build -o bin/$(exec) *.go

clean:
	rm -f bin/$(exec) *~ *#

re: clean all
