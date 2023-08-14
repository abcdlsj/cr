all: clean
	cd ./cmd/ && go generate && cd ..
	gofmt -w -s .

clean:
	rm cr*.go