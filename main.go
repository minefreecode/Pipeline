package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("== Паттерн pipeline представляющйи сосбой обработку данных в несколько этапов ==")
	numbers := generateNumbers(10)
	squared := square(numbers)
	result := addTen(squared)

	for num := range result {
		fmt.Printf("Результат: %d\n", num)
	}
}

func generateNumbers(count int) <-chan int {
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

func square(in <-chan int) <-chan int {
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

func addTen(in <-chan int) <-chan int {
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
