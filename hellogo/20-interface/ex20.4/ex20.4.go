package main

import (
	"github.com/tuckersGo/musthaveGo/ch20/koreaPost"
	"github.com/tuckersGo/musthaveGo/ch20/fedex"
)

// FedexSender, PostSender 구조체 모두 
// 포인터(*FedexSender, *PostSender) 로서  Send() 메소드 구현하고 있음
// 인터페이스 Sender를 통해 위의 Sender 구현체(*FedexSender, *PostSender)를 인자로 넘겨 줌

// Sender 인터페이스 생성
type Sender interface {
	Send(parcel string)
}

// Sender 인터페이스를 입력으로 받음
func SendBook(name string, sender Sender) {
	sender.Send(name)
}

func main() {
	koreaPostSender := &koreaPost.PostSender{}
	SendBook("abc", koreaPostSender)
	SendBook("def", koreaPostSender)

	// Fedex 전송 객체
	fedexSender := &fedex.FedexSender{}
	SendBook("abc123", fedexSender)
	SendBook("def456", fedexSender)
}