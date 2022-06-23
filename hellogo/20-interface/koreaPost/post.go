pacakge koreaPost

import "fmt"

// KoreaPost 제공 패키지 내 전송 담당 구조체
type PostSender struct {
}

func (k *PostSender) Send(parcel string) {
	fmt.Printf("우체국에서 택배 %v를 보냅니다", parcel)
}
