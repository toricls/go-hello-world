FROM public.ecr.aws/docker/library/golang:1.18-buster as builder
WORKDIR /go/src/github.com/toricls/go-hello-world/
RUN go env -w GOPROXY=direct
COPY go.mod go.sum ./
RUN go mod download
COPY main.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM scratch
WORKDIR /root/
COPY --from=builder /go/src/github.com/toricls/go-hello-world/app .
ENTRYPOINT ["./app"]
