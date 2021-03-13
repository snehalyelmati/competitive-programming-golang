package main

import (
	"fmt"
)

type street struct {
	b int
	e int
	n string
	l int
}

type car struct {
	s int
	n []string
}

func main() {
	var d, i, s, v, f int
	fmt.Scanf("%d %d %d %d %d", &d, &i, &s, &v, &f)
	// const size int = s
	streets := make([]street, s)
	for i := 0; i < s; i++ {
		fmt.Scanf("%d %d %s %d", &streets[i].b, &streets[i].e, &streets[i].n, &streets[i].l)
	}
	// fmt.Println(streets)

	cars := make([]car, v)
	for i := 0; i < v; i++ {
		fmt.Scan(&cars[i].s)
		cars[i].n = make([]string, cars[i].s)
		for j := 0; j < cars[i].s; j++ {
			fmt.Scan(&cars[i].n[j])
		}
	}
	fmt.Println(cars)

	streetMap := make(map[string]int)
	for _, car := range cars {
		for _, sName := range car.n {
			_, ok := streetMap[sName]
			if ok {
				streetMap[sName]++
			} else {
				streetMap[sName] = 1
			}
		}
	}
	fmt.Println(streetMap)

	//interMap := make(map[int]string)
	//for _, str := range streets {
	//	_, ok := interMap[str.e]
	//	if ok {
	//		interMap[str.e]++
	//	} else {
	//		interMap[str.e] = 1
	//	}
	//}
	//
	//// fmt.Println()
	//
	//for _, str := range streets {
	//	// fmt.Println(str.n, streetMap[str.n]/str.l)
	//	fmt.Println(str.e)
	//	for k, v := range interMap {
	//		fmt.Println(k)
	//		fmt.Println(stree)
	//		for _, v := range v {
	//			fmt.Println()
	//		}
	//	}
	//}
}