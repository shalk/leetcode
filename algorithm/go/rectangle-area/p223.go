package main

func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
	if A > C {
		A, C = C, A
	}
	if E > G {
		E, G = G, E
	}
	if B > D {
		B, D = D, B
	}
	if F > H {
		F, H = H, F
	}
	// ensure p1 on left of p2 , p3 on left of p4
	l1, d1, r1, u1 := A, B, C, D
	l2, d2, r2, u2 := E, F, G, H
	area1 := area(l1, d1, r1, u1)
	area2 := area(l2, d2, r2, u2)
	length := lap(l1, r1, l2, r2)
	width := lap(d1, u1, d2, u2)
	areaOverlap := length * width
	return area1 + area2 - areaOverlap
}
func area(l1, d1, r1, u1 int) int {
	return (r1 - l1) * (u1 - d1)
}

func lap(l1, r1, l2, r2 int) int {
	if l1 > l2 {
		return lap(l2, r2, l1, r1)
	}
	if r1 <= l2 {
		return 0
	} else if r1 > l2 && r1 <= r2 {
		return r1 - l2
	} else {
		return r2 - l2
	}
}
