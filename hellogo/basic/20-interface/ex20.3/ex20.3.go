package main

import (
	"github.com/tuckersGo/musthaveGo/ch20/koreaPost"
	"github.com/tuckersGo/musthaveGo/ch20/fedex"
)

func SendBook(name string, sender *fedex.FedexSender) {
	sender.Send(name)
}

func main() {
	// koreaPost 전송 객체 생성
	sender := &koreaPost.PostSender{}
	// ERROR! -> ex20.4 인터페이스 구현
	SendBook("abc", sender)
	SendBook("def", sender)
}