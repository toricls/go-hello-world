# go-hello-world

I say hello for you.

```
go run ./main.go
```

```
curl http://localhost
Hello, World!
```

## Customizing message and port number

```
MESSAGE='Hey hey, Yo!' PORT_NUMBER=8080 go run ./main.go
```

```
curl http://localhost:8080
Hey hey, Yo!
```

## Using container

```
docker build -t toricls/go-hello-world:latest .

docker run -p 8000:80 -e MESSAGE='Hello, Container World!' toricls/go-hello-world:latest
```

```
curl http://localhost:8000
Hello, Container World!
```

## Licence

[MIT](LICENSE)
