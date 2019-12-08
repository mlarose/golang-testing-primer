
.PHONY: godoc
godoc:
	docker build -t godoc -f godoc.Dockerfile .
	docker run -itv $(shell cd):/go/src -p 6060:6060 godoc

