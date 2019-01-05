package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}
func findLadders(beginWord string, endWord string, wordList []string) [][]string {

	if len(wordList) == 0 {
		return [][]string{}
	}

	targetIndex := -1
	for i := 0; i < len(wordList); i++ {
		if wordList[i] == endWord {
			targetIndex = i
			break
		}
	}
	if targetIndex == -1 {
		// endWord is not in wordList
		return [][]string{}
	}

	// one trans
	if trans(beginWord, endWord) {
		return [][]string{[]string{beginWord, endWord}}
	}

	// find all short path;
	near := make([][]int, len(wordList))
	for i := 0; i < len(wordList); i++ {
		near[i] = make([]int, 0)
	}

	// O(n^2)
	for i := 0; i < len(wordList); i++ {
		if wordList[i] == beginWord {
			continue
		}
		for j := i + 1; j < len(wordList); j++ {
			if wordList[j] == beginWord {
				continue
			}
			if trans(wordList[i], wordList[j]) {
				near[i] = append(near[i], j)
				near[j] = append(near[j], i)
			}
		}
	}
	// fmt.Println(near)
	paths := make([]Path, 0)
	for i := 0; i < len(wordList); i++ {
		if trans(beginWord, wordList[i]) {
			path := []int{i}
			paths = append(paths, path)
		}
	}

	if len(paths) == 0 {
		return [][]string{}
	}

	// travel by level
	used := map[int]bool{}
	found := false
	for len(paths) > 0 {
		nextPaths := make([]Path, 0)
		// update used
		for _, path := range paths {
			lastIndex := path[len(path)-1]
			used[lastIndex] = true
		}
		for _, path := range paths {
			lastIndex := path[len(path)-1]
			for _, nextIndex := range near[lastIndex] {
				// nextIndex should be not used
				if _, ok := used[nextIndex]; !ok {
					// create newPath and put it int next
					newPath := append(append([]int{}, path...), nextIndex)
					nextPaths = append(nextPaths, newPath)
					if nextIndex == targetIndex {
						found = true
					}
				}
			}
		}
		paths = nextPaths
		if found {
			break
		}
	}

	ret := [][]string{}
	for _, path := range paths {
		lastIndex := path[len(path)-1]
		if lastIndex != targetIndex {
			continue
		}
		words := []string{beginWord}
		for _, index := range path {
			words = append(words, wordList[index])
		}
		ret = append(ret, words)
	}

	return ret
}

type Path []int

func trans(x, y string) bool {
	count := 0
	for i := 0; i < len(x); i++ {
		if x[i] != y[i] {
			count++
		}
	}
	if count == 1 {
		return true
	}
	return false
}
