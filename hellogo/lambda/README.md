


```sh
# Linux 환경
goos=linux go build -o main main.go && zip archive.zip main
# Git bash
goos=linux go build -o main main.go && zip archive.zip main

# Goto AWS lamda console
# Create a function
#  Handler: main(executable)
#  Memory : 128MB
#  Timeout : 
# Action > Upload zip file
# Configure Test event

git archive -o main.zip main.go
```