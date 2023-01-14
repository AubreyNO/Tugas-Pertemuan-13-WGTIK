package main

import "fmt"

func main() {
	var a, b, c, d int
	var p1, c1, p2, c2 int
	fmt.Scan(&a, &b, &c, &d)
	p1 = permutation(a, c)
	c1 = combination(a, c)
	p2 = permutation(b, d)
	c2 = combination(b, d)
	fmt.Println(p1, c1)
	fmt.Println(p2, c2)
}

func findFactorial(n int, result *int) {
	for n > 0 {
		*result = (*result) * n
		n--
	}
}

func permutation(n, r int) int {
	var fakt1, fakt2 int
	fakt1 = 1
	fakt2 = 1
	findFactorial(n, &fakt1)
	findFactorial((n - r), &fakt2)

	return fakt1 / fakt2
}

func combination(n, r int) int {
	var fakt1, fakt2, fakt3 int
	fakt1 = 1
	fakt2 = 1
	fakt3 = 1
	findFactorial(n, &fakt1)
	findFactorial(r, &fakt2)
	findFactorial((n - r), &fakt3)
	return fakt1 / (fakt2 * fakt3)
}
