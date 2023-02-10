











- copies bytes between interfaces and treat data like a stream
- pipe read from a stream and write from the same stream to a different source

```sh
cd hellogo/cookbook
mkdir -p chapter1/interfaces && cd $_
go mod init github.com/jnuho/jnuho.github.io/hellogo/cookbook/chapter1/interfaces
go mod tidy
# create .go files
cd chapter1/interfaces/example
go build
cat go.mod
  module github.com/jnuho/jnuho.github.io/hellogo/cookbook
  go 1.18
```

- Implement `io.Writer` and `io.Reader`
  - io.Copy() 구현
  - io.Pipe() 활용

```
ㄴcookbook
  ㄴinterfaces
      interfaces.go
    pipes.go
    ㄴexample
      main.go
```

- bytes and strings packages to convert the data from string to byte types

