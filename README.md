# urip
urip server returns your ip address over http.

## Usage

When using port 80, the -a option can be omitted.
```bash
urip -a hostname:port # e.g. urip -a :8080
```

## Installation

```bash
go get github.com/minami14.com/urip
```

## Docker
```bash
docker run -p 80:80 minami14/urip
```
