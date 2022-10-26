```sh
mkdir src
cd src
export GOFLAGS=-mod=vendor
export GO111MODULE=on
go mod init github.com/jnuho.github.io/hellogo/go-web-docker

mkdir conf views
vim main.go
```



```
HTTP-SERVER
	cmd/http-server
		main.go
	pkg/server
		http.go
	Dockerfile
	go.mod
	stack.yaml
```



go run main.go -port 8080

curl -X POST -H "Content-Type: application/json" \
    -d '{"name": "linuxize", "email": "linuxize@example.com"}' \
    http://localhost:8080/targetgroup


