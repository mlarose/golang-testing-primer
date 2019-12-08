
.PHONY: test
test:
	go test -count=1 -v ./...

.PHONY: godoc
godoc:
	docker build -t godoc -f godoc.Dockerfile .
	docker run -itv $(shell pwd):/go/src -p 6060:6060 godoc

.PHONY: bench
bench:
	go test -bench . github.com/mlarose/golang-testing-tutorial/...
