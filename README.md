### Description

This project is a simple HTTP server that prints out the headers and/or bodies of incoming requests to stdout.

### Build

```
go build -o ./bin/httparrot .
```

### Run

```
./bin/httparrot -port=5555 -header=true -body=true
```

### Docker Build

```
docker build -t httparrot .
```

### Docker Run

```
docker run -it --rm -p 5555:5555 --name httparrot-running httparrot
```

### Test

```
curl -XPOST -H "Content-Type: application/json" localhost:5555/anyroute -d '{"hello", "world"}'
```
