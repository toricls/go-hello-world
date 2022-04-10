# go-hello-world

I say hello for you.

```shell
go run ./main.go
```

```shell
curl http://localhost
Hello, World!
```

## Usage

### Customize message and port number

```shell
MESSAGE='Hey hey, Yo!' PORT_NUMBER=8080 SLEEP_MILLISEC=300 go run ./main.go
```

```shell
curl http://localhost:8080
Hey hey, Yo!
```

## Use pre-built container image

```shell
docker run -p 8000:80 -e MESSAGE='Hello, Container World!' public.ecr.aws/toricls/go-hello-world:latest
```

## Build and use your own container image

```shell
docker build -t go-hello-world:latest .

docker run -p 8000:80 -e MESSAGE='Hello, Container World!' go-hello-world:latest
```

## Run on AWS Lambda

`go-hello-world` natively supports running on AWS Lambda as container image. Try deploying a new AWS Lambda function in your AWS account using the pre-build container image `public.ecr.aws/toricls/go-hello-world:latest` :)

See the `main` function in [`main.go`](main.go) to know how you can implement in the same way in your go app.

Note that the `PORT_NUMBER` environment variable is not (obviously) supported when you run the app on AWS Lambda.

## Contribution

1. Fork ([https://github.com/toricls/go-hello-world/fork](https://github.com/toricls/go-hello-world/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Create a new Pull Request

## Licence

[MIT](LICENSE)

## Author

[Tori](https://github.com/toricls)
