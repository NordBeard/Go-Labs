package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
	"sync"
	"time"
)

// Матрица для свёртки (Гауссово размытие)
var kernel = [3][3]float64{
	{1, 2, 1},
	{2, 4, 2},
	{1, 2, 1},
}

// applyKernel применяет свертку с матрицей ядра на каждый пиксель
func applyKernel(img *image.RGBA, result *image.RGBA, x, y int) {
	var r, g, b float64
	bounds := img.Bounds()

	// Применяем фильтр для пикселя с учётом соседей
	for ky := -1; ky <= 1; ky++ {
		for kx := -1; kx <= 1; kx++ {
			px, py := x+kx, y+ky
			if px >= bounds.Min.X && px < bounds.Max.X && py >= bounds.Min.Y && py < bounds.Max.Y {
				original := img.RGBAAt(px, py)
				weight := kernel[ky+1][kx+1]
				r += float64(original.R) * weight
				g += float64(original.G) * weight
				b += float64(original.B) * weight
			}
		}
	}

	// Нормализуем результат и ограничиваем его диапазон
	weightSum := 16.0
	r = math.Min(r/weightSum, 255)
	g = math.Min(g/weightSum, 255)
	b = math.Min(b/weightSum, 255)

	// Записываем результат в новый пиксель
	result.SetRGBA(x, y, color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: 255})
}

// filterWithKernel обрабатывает строку пикселей с фильтром
func filterWithKernel(img *image.RGBA, result *image.RGBA, y int, wg *sync.WaitGroup) {
	defer wg.Done()
	bounds := img.Bounds()
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		applyKernel(img, result, x, y)
	}
}

func main() {
	// Открытие изображения
	inputFile, err := os.Open("C:/For study/Python Projects/5_lab/1_task/2_task/3_task/4_task/7a941b62e881e1174c6213982305a636.png")
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

	// Преобразуем изображение в *image.RGBA
	rgbaImg, ok := img.(*image.RGBA)
	if !ok {
		fmt.Println("Изображение не в формате RGBA")
		return
	}

	// Параллельная обработка с использованием фильтра
	start := time.Now()
	var wg sync.WaitGroup
	bounds := rgbaImg.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		wg.Add(1)
		go filterWithKernel(rgbaImg, rgbaImg, y, &wg)
	}
	wg.Wait()
	fmt.Printf("Параллельная обработка заняла: %v\n", time.Since(start))

	// Сохранение результата
	outputFile, err := os.Create("C:/photolab/4.png")
	if err != nil {
		fmt.Println("Ошибка создания файла:", err)
		return
	}
	defer outputFile.Close()

	if err := png.Encode(outputFile, rgbaImg); err != nil {
		fmt.Println("Ошибка сохранения изображения:", err)
	}
	fmt.Println("Задание 4 завершено.")
}
