all: go

exec = wordchains

go: main.go
	go test ./wordchains/
	go build -o bin/$(exec) main.go

clean:
	rm -f bin/$(exec) *~ *#
	rm -rf src/gopkg.in

re: clean all
