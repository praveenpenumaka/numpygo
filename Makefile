test:
	go test ./...

test-coverage:
	go test -covermode=count -coverprofile=count.out ./...
	go tool cover -html=count.out

clean:
	rm -i count.out coverage.out