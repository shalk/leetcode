package main

import "fmt"

/**
 * Definition for a point.
 * type Point struct {
 *     X int
 *     Y int
 * }
 */
func maxPoints(points []Point) int {
	if len(points) == 0 {
		return 0
	}

	m := make([]int, len(points))
	for i := 0; i < len(points); i++ {
		if m[i] == -1 {
			continue
		}
		m[i] = 1
		for j := i + 1; j < len(points); j++ {
			if equal(points[i], points[j]) {
				m[i]++
				m[j] = -1
			}
		}
	}
	max := 0
	visited := make([][]bool, len(points))
	for i := 0; i < len(points); i++ {
		visited[i] = make([]bool, len(points))

		if m[i] > max {
			max = m[i]
		}
	}

	for i := 0; i < len(points); i++ {
		if m[i] == -1 {
			continue
		}

		for j := i + 1; j < len(points); j++ {
			if m[j] == -1 {
				continue
			}
			if visited[i][j] {
				continue
			}
			count := m[i] + m[j]

			visited[i][j] = true
			for k := j + 1; k < len(points); k++ {
				if m[k] == -1 {
					continue
				}

				if check(points[i], points[j], points[k]) {
					count += m[k]
					visited[i][k] = true
					visited[j][k] = true
				}
			}

			if count > max {
				max = count
			}
		}

	}
	return max
}

func equal(a, b Point) bool {
	return a.X == b.X && a.Y == b.Y
}
func check(a, b, c Point) bool {
	if equal(a, b) || equal(b, c) || equal(a, c) {
		return true
	}
	if a.Y == b.Y {
		return c.Y == a.Y
	} else {
		if a.Y == c.Y {
			return false
		} else {
			return (a.X-b.X)*(a.Y-c.Y) == (a.X-c.X)*(a.Y-b.Y)
		}
	}

}
