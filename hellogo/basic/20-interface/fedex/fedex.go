package fedex

import "fmt"

// Fedex제공 패키지 내 전송 담당 구조체
type FedexSender struct {
}

func (f *FedexSender) Send(parcel string) {
	fmt.Printf("Fedex sends %v parcel\n", parcel)
}
