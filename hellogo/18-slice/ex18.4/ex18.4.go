package main
import "fmt"

// ERROR : func changeArray(array2 *[5]int) {
// array 포인터 사용 못함, 대신에 slice 사용
// array를 매개변수로 사용 시 배열 전체가 복사됨 (비효율적)
func changeArray(array2 [5]int) {
	array2[2] = 200
}

// slice2는 []int 타입 (슬라이스)이며
// 포인터(8바이트 메모리주소), len(8-byte), cap(8-byte) 세개의 필드를 갖는 구조체임
// 가리키는 배열 크기 상관없이 24바이트 복사됨
func changeSlice(slice2 []int) {
	slice2[2] = 200
}

func main() {
	array := [5]int{1,2,3,4,5}
	slice := []int{1,2,3,4,5}

	changeArray(array)
	changeSlice(slice)

	fmt.Println("array: ", array)
	fmt.Println("slice: ", slice)
}