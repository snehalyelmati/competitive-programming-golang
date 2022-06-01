package main

import "fmt"

func permutations(s string) []string {
	return helper("", s)
}

func helper(p, u string) []string {
	if u == "" {
		return []string{p}
	}

	res := []string{}
	for i := 0; i < len(p); i++ {
		fmt.Println(string(p[:i])+string(u[0])+string(p[i:]), u[1:])
		res = append(res, helper(string(p[:i])+string(u[0])+string(p[i:]), u[1:])...)
	}
	res = append(res, helper(p+string(u[0]), u[1:])...)
	return res
}

func main() {
	s := "abcd"
	fmt.Println(permutations(s))
}
