package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"sync"
	"time"
)

// Функция для обработки строки пикселей в параллельном режиме
func processRow(img *image.RGBA, y int, wg *sync.WaitGroup) {
	defer wg.Done()
	bounds := img.Bounds()
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		original := img.RGBAAt(x, y)
		gray := uint8((uint16(original.R) + uint16(original.G) + uint16(original.B)) / 3)
		img.SetRGBA(x, y, color.RGBA{R: gray, G: gray, B: gray, A: original.A})
	}
}

func main() {
	// Открытие изображения
	inputFile, err := os.Open("C:/For study/Python Projects/5_lab/1_task/2_task/3_task/7a941b62e881e1174c6213982305a636.png")
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		return
	}
	defer inputFile.Close()

	// Декодирование изображения
	img, _, err := image.Decode(inputFile)
	if err != nil {
		fmt.Println("Ошибка декодирования изображения:", err)
		return
	}

	// Преобразование к *image.RGBA (для обработки в дальнейшем)
	rgbaImg, ok := img.(*image.RGBA)
	if !ok {
		fmt.Println("Изображение не в формате RGBA")
		return
	}

	// Параллельная обработка
	start := time.Now()
	var wg sync.WaitGroup
	bounds := rgbaImg.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		wg.Add(1)
		go processRow(rgbaImg, y, &wg)
	}
	wg.Wait()
	fmt.Printf("Параллельная обработка заняла: %v\n", time.Since(start))

	// Сохранение результата
	outputFile, err := os.Create("C:/photolab/3.png")
	if err != nil {
		fmt.Println("Ошибка создания файла:", err)
		return
	}
	defer outputFile.Close()

	if err := png.Encode(outputFile, rgbaImg); err != nil {
		fmt.Println("Ошибка сохранения изображения:", err)
	}
	fmt.Println("Задание 3 завершено.")
}
