package functions

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	s_slice := strings.Fields(s)
	answer := make(map[string]int)

	for i := 0; i < len(s_slice); i++ {
		fmt.Println(s_slice[i])
		if answer[s_slice[i]] == 0 {
			answer[s_slice[i]] = 1
		} else {
			answer[s_slice[i]] += 1
		}
	}
	return answer
}

func Exercise() {
	WordCount("I ate a donut. Then I ate another donut.")
}
