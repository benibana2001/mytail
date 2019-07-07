run:
	go run main.go sample.txt

test:
	go test ./tail

fmt:
	go fmt main.go
	go fmt ./tail/tail.go
	go fmt ./tail/tail_test.go
