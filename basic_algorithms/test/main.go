package main

import (
	"fmt"
)

func filterCharacter(s, target string) string {
	if s < target || target == "" {
		return s
	}

	if s[:len(target)] == target {
		return filterCharacter(s[len(target):], target)
	}
	return s[:len(target)] + filterCharacter(s[len(target):], target)
}

func subSequence(p, u string) []string {
	if u == "" {
		fmt.Println(p)
		return []string{p}
	}

	left := subSequence(p+string(u[0]), u[1:])
	right := subSequence(p, u[1:])
	return append(left, right...)
}

func countMazePaths(maze [][]int, tr, tc int) int {
	return helperMaze(0, 0, tr, tc)
}

func helperMaze(r, c, tr, tc int) int {
	if r == tr-1 && c == tc-1 {
		return 1
	} else if r >= tr || c >= tc {
		return 0
	}

	return helperMaze(r, c+1, tr, tc) + helperMaze(r+1, c, tr, tc)
}

func printMazePaths(maze [][]int, tr, tc int) []string {
	return helperMazePath(maze, "", 0, 0, tr, tc)
}

func helperMazePath(maze [][]int, path string, r, c, tr, tc int) []string {
	if maze[r][c] == 1 {
		return []string{path}
	}

	paths := []string{}

	if c+1 < tc {
		fmt.Println(path, r, c, tr, tc)
		paths = append(paths, helperMazePath(maze, path+"R", r, c+1, tr, tc)...)
	}
	if r+1 < tr {
		fmt.Println(path, r, c, tr, tc)
		paths = append(paths, helperMazePath(maze, path+"D", r+1, c, tr, tc)...)
	}

	return paths
}

func printMazePathsObstacles(maze [][]int, tr, tc int) []string {
	return helperMazePathObstacles(maze, "", 0, 0, tr, tc)
}

func helperMazePathObstacles(maze [][]int, path string, r, c, tr, tc int) []string {
	// fmt.Println(path, r, c, tr, tc, maze[r][c])
	if maze[r][c] == 1 {
		return []string{path}
	}

	if maze[r][c] == -1 {
		return []string{}
	}

	maze[r][c] = -1
	paths := []string{}

	if c+1 < tc {
		paths = append(paths, helperMazePathObstacles(maze, path+"R", r, c+1, tr, tc)...)
	}
	if r+1 < tr {
		paths = append(paths, helperMazePathObstacles(maze, path+"D", r+1, c, tr, tc)...)
	}
	if c-1 > 0 {
		paths = append(paths, helperMazePathObstacles(maze, path+"L", r, c-1, tr, tc)...)
	}
	if r-1 > 0 {
		paths = append(paths, helperMazePathObstacles(maze, path+"U", r-1, c, tr, tc)...)
	}

	return paths
}

func main() {
	// s, target := "applebananapple", ""
	// fmt.Println(filterCharacter(s, target))

	// sub := "abc"
	// subSeq := subSequence("", sub)
	// sort.Strings(subSeq)
	// fmt.Println(subSeq)

	maze := [][]int{
		{0, 0, -1, -1},
		{0, -1, 0, 1},
		{0, -1, 0, -1},
		{0, 0, 0, 0},
	}
	for i := 0; i < len(maze); i++ {
		fmt.Println(maze[i])
	}
	targetRow, targetColumn := len(maze), len(maze)
	// fmt.Println(countMazePaths(maze, targetRow, targetColumn))
	// fmt.Println(printMazePaths(maze, targetRow, targetColumn))
	fmt.Println(printMazePathsObstacles(maze, targetRow, targetColumn))
}
