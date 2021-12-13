package main

import "fmt"

func isValid(s string) bool {
	arr := make([]byte, 0)
	arr = append(arr, s[0])

	bracketMap := map[byte]byte{')': '(', ']': '[', '}': '{'}

	for i := 1; i < len(s); i++ {
		// fmt.Println(string(s[i]), arr, len(s))
		if len(arr) > 0 && bracketMap[s[i]] == arr[len(arr)-1] {
			arr = arr[:len(arr)-1]
		} else {
			arr = append(arr, s[i])
		}
	}
	// fmt.Println(arr)

	return len(arr) == 0
}

func getOpp(b string) string {
	if b == ")" {
		return "("
	} else if b == "]" {
		return "["
	} else if b == "}" {
		return "{"
	}
	return "-1"
}

func main() {
	s := "{()}[]"
	fmt.Println(isValid(s))
}
