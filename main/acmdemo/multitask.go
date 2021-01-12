package acmdemo

import "fmt"

type FactorialResult struct {
	N, Result int
}

func factorial(n int, m chan FactorialResult) {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	m <- FactorialResult{N: n, Result: result}
}

func MainMultitask() {
	fmt.Println("==============================================================")
	fmt.Println("Multitask")
	fmt.Println("==============================================================")

	var channels []chan FactorialResult
	for _, v := range []int{33, 2, 12} {
		ch := make(chan FactorialResult)
		channels = append(channels, ch)
		go factorial(v, ch)
	}

	for _, v := range channels {
		fmt.Println(<-v)
	}
}
