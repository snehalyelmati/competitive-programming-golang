package main

import "fmt"

func productExceptSelf(nums []int) []int {
	N := len(nums)
	res := make([]int, N)

	res[N-1] = 1
	for i := N-2; i>=0; i-- {
		res[i] = res[i+1]*nums[i+1]
	}

	leftProd := 1
	for i := 0; i < N-1; i++ {
		leftProd *= nums[i]
		res[i+1] *= leftProd
	}

	return res
}
	
func productExceptSelfOld2(nums []int) []int {
	res := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		temp := 1
		for j := 0; j < len(nums); j++ {
			if j != i {
				temp *= nums[j]
			}
		}
		res = append(res, temp)
	}

	return res
}
	
func productExceptSelfOld(nums []int) []int {
	prod := 1
	zero := 0
    for _, v := range nums {
		if v == 0 {
			zero++
		} else {
			prod *= v
		}
	}

	res := make([]int, 0)
	for _, v := range nums {
		if zero >= 2 {
			res = append(res, 0)
		} else if zero == 1 {
			if v == 0 {
				res = append(res, prod)
			} else {
				res = append(res, 0)
			}
		} else {
			res = append(res, prod/v)
		}
	}

	return res
}

func main() {
	nums := []int {-1, 1, 0, 3, -3}
	// nums := []int {1, 2, 3, 4}
	fmt.Println(nums)
	fmt.Println(productExceptSelf(nums))
}