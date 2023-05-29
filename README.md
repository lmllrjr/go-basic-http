# go-basic-http
[![Test](https://github.com/lmllrjr/go-basic-http/actions/workflows/test.yaml/badge.svg)](https://github.com/lmllrjr/go-basic-http/actions/workflows/test.yaml)

## run webworker
```sh
go mod tidy
go run ./cmd/main.go
```

## make some curl requests
```sh
curl localhost:8080 -u 007
```
>**Note**:  
>The root path `/` (hello world REST API) is the only one with a full trace.

```sh
curl localhost:8080/greet/luke%20skywalker -u 007
```

```sh
curl localhost:8080/slug/777 -u 007
```

>**Note**:  
>The password is `123` for basic auth.

## run tests
```sh
go test ./...
```
