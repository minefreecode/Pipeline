package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("== Паттерн pipeline представляющий сосбой обработку данных в несколько этапов ==")
	numbers := generate(10)
	squared := multiplySquare(numbers)
	result := sumTen(squared)

	for num := range result {
		fmt.Printf("Результат: %d\n", num)
	}
}

func generate(count int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := 0; i < count; i++ {
			num := rand.Intn(10) + 1
			fmt.Printf("Сгенерировано: %d\n", num)
			out <- num
			time.Sleep(100 * time.Millisecond)
		}
	}()
	return out
}

func multiplySquare(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for num := range in {
			squared := num * num
			fmt.Printf("Squared %d -> %d\n", num, squared)
			out <- squared
			time.Sleep(150 * time.Millisecond)
		}
	}()
	return out
}

func sumTen(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for num := range in {
			result := num + 10
			fmt.Printf("Добавлено 10 к %d -> %d\n", num, result)
			out <- result
			time.Sleep(100 * time.Millisecond)
		}
	}()
	return out
}
