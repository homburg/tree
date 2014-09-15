benchmark:
	go test -bench .

time:
	cat data.csv | time -p tree , > /dev/null

test:
	go test -v ./...

travis_install:
	go get -d -v ./... && go build -v ./...

travis: travis_install test benchmark

