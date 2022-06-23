package main
type Attacker interface {
	Attack()
}

func main() {
	var att Attacker // 기본값은 nil 임
	att.Attack() //nil이 함수 호출 시도 해서 ERROR 발생
}