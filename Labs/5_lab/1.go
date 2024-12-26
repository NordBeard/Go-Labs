package main

import (
	"fmt"
	"time"
)

// Функция count считывает числа из канала и выводит их квадрат
func count(ch <-chan int) {
	for num := range ch { // Чтение из канала до его закрытия
		fmt.Printf("Квадрат числа %d: %d\n", num, num*num)
	}
}

func main() {
	// Создаём канал для передачи чисел
	ch := make(chan int)

	// Запускаем функцию count в отдельной горутине
	go count(ch)

	// Отправляем числа в канал
	for i := 1; i <= 5; i++ {
		ch <- i
	}
	// Закрываем канал
	close(ch)

	// Ждём завершения работы горутины
	time.Sleep(time.Second)
	fmt.Println("Задание 1 завершено.")
}
