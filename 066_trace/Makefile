
run:
	echo ${GOVERSION}
	docker build -t test --build-arg GOVERSION=${GOVERSION} .
	rm -rf trace.out
	docker run --rm --name test -v $$PWD/:/workspace test \
		/bin/bash -c \
		'go version; GOMAXPROCS=2 go run main.go'
	go tool trace  trace.out
