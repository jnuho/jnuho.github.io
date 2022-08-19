package main

import (
	"github.com/tuckersGo/musthaveGo/ch20/fedex"
)

func SendBook(name string, sender *fedex.FedexSender) {
	sender.Send(name)
}

/*
cd 20-interface/ex20.2
go mod init 20-interface/ex20.2
go mod tidy
*/
func main() {
	// Fedex 전송 객체 생성
	sender := &fedex.FedexSender{}
	SendBook("어린 왕자", sender)
	SendBook("그리스인 조르바", sender)
}