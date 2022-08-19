package ducktape

import "fmt"

type Database interface {
	Get()
	Set()
}

// ducktape: Database 구현한다는걸 명시할 필요 없음
type BDatabase struct {
}

// ducktape: Database 구현한다는걸 명시할 필요 없음
type CDatabase struct {
}

func TotalTime(db Database) int {
	db.Get()
	db.Set()
}

func Compare(){
	BDB := &BDatabase{}
	CDB := &CDatabase{}

	if TotalTime(BDB) < TotalTime(CDB) {
		fmt.Println("B회사 제품이 더 빠릅니다.")
	} else {
		fmt.Println("C회사 제품이 더 빠릅니다.")
	}
}
