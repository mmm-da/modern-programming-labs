package main

import "math"

type task1 struct {
	foo func(a int) bool
}

type task2 struct {
	foo func(a, b string) string
}

type task3 struct {
	foo func(a, b, c float64) float64
}

func main() {
	t1 := task1{}
	t1.foo = func(a int) bool {
		return a%13 == 0
	}

	t2 := task2{}
	t2.foo = func(a, b string) string {
		if len(a) >= len(b) {
			return a
		}
		return b
	}

	t3 := task3{}
	t3.foo = func(a, b, c float64) float64 {
		return b*b - 4*a*c
	}

	t4 := task3{}
	t4.foo = func(a, b, c float64) float64 {
		return a * math.Pow(b, c)
	}
}
