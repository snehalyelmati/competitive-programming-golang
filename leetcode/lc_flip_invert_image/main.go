package main

import "fmt"

func flipAndInvertImage(image [][]int) [][]int {
	// flip
	if len(image) > 1 {
		for i := 0; i < len(image); i++ {
			for j := 0; j < len(image[i])/2; j++ {
				image[i][j], image[i][len(image[i])-j-1] = image[i][len(image[i])-j-1], image[i][j]
			}
		}
	}
	fmt.Println(image)

	// invert
	for i := 0; i < len(image); i++ {
		for j := 0; j < len(image[i]); j++ {
			image[i][j] ^= 1
			// image[i][j] = (image[i][j] + 1) % 2
		}
	}

	return image
}

func main() {
	image := [][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}}
	fmt.Println(image)
	fmt.Println(flipAndInvertImage(image))
}
