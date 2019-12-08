FROM golang

EXPOSE 6060

RUN go get golang.org/x/tools/cmd/godoc

CMD ["godoc", "-http=:6060"]