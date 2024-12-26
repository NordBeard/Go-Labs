package main

import (
	"math/rand"
	"strconv"
	"syscall/js"
	"time"
)

type Point struct {
	x, y int
}

const (
	boardWidth  = 20
	boardHeight = 20
	cellSize    = 20
)

var (
	snake     []Point
	direction Point
	food      Point
	gameOver  bool
	score     int
	ctx       js.Value
)

func initGame() {
	snake = []Point{{boardWidth / 2, boardHeight / 2}}
	direction = Point{1, 0}
	placeFood()
	score = 0
	gameOver = false
}

func placeFood() {
	food = Point{rand.Intn(boardWidth), rand.Intn(boardHeight)}
}

func drawBoard() {
	// Очистка экрана
	ctx.Call("clearRect", 0, 0, boardWidth*cellSize, boardHeight*cellSize)

	// Рисуем еду
	ctx.Set("fillStyle", "red")
	ctx.Call("fillRect", food.x*cellSize, food.y*cellSize, cellSize, cellSize)

	// Рисуем змейку
	ctx.Set("fillStyle", "green")
	for _, segment := range snake {
		ctx.Call("fillRect", segment.x*cellSize, segment.y*cellSize, cellSize, cellSize)
	}

	// Рисуем счет
	ctx.Set("fillStyle", "black")
	ctx.Call("fillText", "Score: "+strconv.Itoa(score), 10, boardHeight*cellSize+20)
}

func moveSnake() {
	head := snake[len(snake)-1]
	newHead := Point{head.x + direction.x, head.y + direction.y}

	// Проверка столкновений
	if newHead.x < 0 || newHead.x >= boardWidth || newHead.y < 0 || newHead.y >= boardHeight {
		gameOver = true
		return
	}
	for _, segment := range snake {
		if segment == newHead {
			gameOver = true
			return
		}
	}

	snake = append(snake, newHead)

	// Проверка еды
	if newHead == food {
		score++
		placeFood()
	} else {
		snake = snake[1:] // Убираем хвост
	}
}

func handleKeyDown(this js.Value, args []js.Value) interface{} {
	key := args[0].Get("keyCode").Int()
	switch key {
	case 37: // Left
		if direction != (Point{1, 0}) {
			direction = Point{-1, 0}
		}
	case 38: // Up
		if direction != (Point{0, 1}) {
			direction = Point{0, -1}
		}
	case 39: // Right
		if direction != (Point{-1, 0}) {
			direction = Point{1, 0}
		}
	case 40: // Down
		if direction != (Point{0, -1}) {
			direction = Point{0, 1}
		}
	}
	return nil
}

func gameLoop() {
	if gameOver {
		ctx.Set("fillStyle", "black")
		ctx.Call("fillText", "Game Over! Refresh to restart.", 10, boardHeight*cellSize/2)
		return
	}

	moveSnake()
	drawBoard()

	time.AfterFunc(200*time.Millisecond, gameLoop)
}

func main() {
	// Получаем доступ к контексту рисования на холсте
	doc := js.Global().Get("document")
	canvas := doc.Call("getElementById", "gameCanvas")
	ctx = canvas.Call("getContext", "2d")

	// Инициализация игры
	initGame()

	// Добавляем обработчик событий клавиатуры
	js.Global().Call("addEventListener", "keydown", js.FuncOf(handleKeyDown))

	// Запускаем игровой цикл
	gameLoop()

	// Не позволяем завершить программу
	select {}
}
