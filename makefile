all: clean
	cd ./cmd/ && go generate && cd ..
	gofmt -w -s .

clean:
	rm cr*.go > /dev/null 2>&1 || true