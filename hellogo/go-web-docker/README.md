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
MathApp
    ├── Dockerfile
    ├── Dockerfile.production
    └── src
        ├── conf
        │   └── app.conf
        ├── go.mod
        ├── go.src
        ├── main.go
        ├── main_test.go
        ├── vendor
        └── views
            ├── invalid-route.html
            └── result.html
```
