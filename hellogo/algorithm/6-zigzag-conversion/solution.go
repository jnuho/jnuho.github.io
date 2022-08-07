package main

import "fmt"

// 0     0     0
// 1    91    9
// 2   8 2   8
// 3  7  3  7
// 4 6   4 6
// 5     5
// 2*(numRows -i-1)

// 0    8
// 1   79   5
// 2  6 0  4
// 3 5  1 3
// 4    2

// 0 1 2 3 4 5 6 7 8 9 0 1 2 3
// P A Y P A L I S H I R I N G

// 0   6   2
// 1  57  13
// 2 4 8 0
// 3   9

func convert(s string, numRows int) string {
	ii := 0
	res := ""
	ll := len(s)

	if ll < 2 {
		return s
	}

	cnt := 0
	
	for i := range s[: numRows] {
		cnt = 0
		ii = i
		if ii == 0 {
			for ii < ll {
				res += string(s[ii])
				ii += 2*(numRows -i-1)
			}
		} else if ii*2 < numRows {
			for ii < ll {
				res += string(s[ii])
				if cnt % 2 == 0 {
					ii += 2*(numRows -i-1)
				} else {
					ii += 2*i
				}
				cnt++
			}
		} else if ii*2 >= numRows && ii < numRows-1 {
			for ii < ll {
				res += string(s[ii])
				if cnt % 2 == 0 {
					ii += 2*(numRows -i-1)
				} else {
					ii += 2*i
				}
				cnt++
			}
		} else if ii == numRows-1 {
			for ii < ll{
				res += string(s[ii])
				ii += 2*i
			}
		}
	}

	return res
}


func main() {
	fmt.Println(convert("A", 1))
	fmt.Println(convert("PAYPALISHIRING", 3))
	fmt.Println(convert("PAYPALISHIRING", 4))
}