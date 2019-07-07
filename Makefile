run:
	go run main.go sample.txt

test:
	go test ./tail

fmt:
	go fmt main.go ./tail.go
