1. There are 3 tasks in the test. You can solve them in any order.

2. There's no option to pause. Make sure you will not be interrupted for 120 minutes.

3. Your solution(s) should consider all possible corner cases and handle large input efficiently. Passing the example test does not indicate that your solution is correct. The example test is not part of your final score.

4. After finishing the test you will receive feedback containing your score. See example feedback.

5. If you accidentally close your browser, use the invitation link to get back to your test.

6. Hint: you can use your own IDE and use copy-paste, but make sure your solution compiles in Codility's environment.

* You can write your solution(s) in C, C++, Go, Java 11, Java 8 or Kotlin.



# 1

```go

package solution

// you can also use imports, for example:
//import "fmt"
import "strings"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func getAlphabets() []string {
    alphabet := make([]string, 0)
    var ch byte
    for ch = 'a'; ch <= 'z'; ch++ {
        alphabet = append(alphabet, string(ch))
    }
    return alphabet
}

/**
string내, 알파벳 반복횟수 계산하기
결과가 -1 이면 prime number => 'a' N만큼반복 하여 리턴
*/
func findDivNum(N int) int {
    res := N
    for i:=2; i<=N; i++{
        if N%i == 0{
            if i<=res{
                res=i
            }
        }
    }
    if res == N {
        return -1
    } else{
        return res
    }
}

func Solution(N int) string {
    // write your code in Go 1.13.8

    // ['a', .., 'z']
    alphabets := getAlphabets()

    if N <= 26 {
        return strings.Join(alphabets[:N], "")
    } else {
        div := findDivNum(N)
        if div == -1 {
            return strings.Repeat("a", N)
        } else {
            return ""
        }
    }
}
```

# 2

```go
package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(S string, A []int) string {
    // write your code in Go 1.13.8
    if len(S) == 0{
        return ""
    }

    idx := 0
    res := string(S[0])
    for {
        idx = A[idx]
        if idx == 0 {
            break
        }
        res += string(S[idx])
    }
    return res
}


```


# 3

```go
package solution

// you can also use imports, for example:
//import "fmt"
// import "os"
import (
    "strings"
    "strconv"
)

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(S string, C string) int {
    rows := strings.Split(S, "\n")
    colNames := strings.Split(rows[0], ",")
    dataRows := rows[1:]

    idx := 0
    for i:=0; i<len(colNames); i++ {
        if C == colNames[i] {
            idx = i
        }
    }

    maxVal := -10000
    for i:=0; i<len(dataRows); i++ {
        temp, _ :=  strconv.Atoi(strings.Split(dataRows[i], ",")[idx])
        if maxVal <= temp {
            maxVal = temp
        }
    }
    return maxVal
    // write your code in Go 1.13.8
}
```
