FROM public.ecr.aws/bitnami/golang:1.16
WORKDIR /go/src/github.com/toricls/go-hello-world/
COPY go.mod .
COPY main.go .
RUN go env -w GOPROXY=direct
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM scratch
WORKDIR /root/
COPY --from=0 /go/src/github.com/toricls/go-hello-world/app .
CMD ["./app"]
