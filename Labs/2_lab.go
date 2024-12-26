package main

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

// 1.1
func formatIP(ip [4]byte) (string, error) {
	var strIP []string
	for _, i := range ip {
		strIP = append(strIP, strconv.Itoa(int(i)))
	}

	return strings.Join(strIP, "."), nil
}

// 1.2
func listEven(a, b int) ([]int64, error) {
	size := b - a + 1
	var evenNumber []int64
	if size <= 0 {
		return evenNumber, errors.New("index out of range")
	}

	for i := a; i <= b; i++ {
		if i%2 == 0 {
			evenNumber = append(evenNumber, int64(i))
		}
	}

	return evenNumber, nil
}

// 2
func countCharacters(s string) (map[rune]int, error) {
	counts := make(map[rune]int)

	for _, char := range s {
		counts[char]++
	}

	return counts, nil
}

// 3
type Dot struct {
	x, y float64
}

type segment struct {
	dots [2]Dot
}

func (s *segment) length() float64 {
	dx := float64(s.dots[1].x - s.dots[0].x)
	dy := float64(s.dots[1].y - s.dots[0].y)
	return math.Sqrt(dx*dx + dy*dy)
}

func NewSegment(dots [2]Dot) (segment, error) {
	newSegment := segment{dots: dots}
	if newSegment.length() == 0 {
		return segment{}, errors.New("это не отрезок")
	}
	return newSegment, nil
}

type Shape interface {
	Area() float64
}

type triangle struct {
	dots [3]Dot
}

func (t *triangle) Area() float64 {
	p1p2, err := NewSegment([2]Dot{t.dots[0], t.dots[1]})
	p2p3, err1 := NewSegment([2]Dot{t.dots[1], t.dots[2]})
	p1p3, err2 := NewSegment([2]Dot{t.dots[0], t.dots[2]})
	if err != nil || err1 != nil || err2 != nil {
		fmt.Println("Ошибка: ", err, err1, err2)
		return 0
	}

	p := (p1p2.length() + p2p3.length() + p1p3.length()) / 2
	area := math.Sqrt(p * (p - p1p2.length()) * (p - p2p3.length()) * (p - p1p3.length()))

	return area
}

func NewTriangle(d [3]Dot) (triangle, error) {
	newTriangle := triangle{dots: d}
	area := newTriangle.Area()
	if area == 0 {
		return triangle{}, errors.New("это не треугольник")
	}
	return newTriangle, nil
}

type circle struct {
	centreDot Dot
	radius    float64
}

func (c *circle) Area() float64 {
	area := 3.14 * (math.Pow(c.radius, 2))
	return area
}

func NewCircle(centerDot Dot, radius float64) (circle, error) {
	if radius == 0 {
		return circle{}, errors.New("радиус не может быть равен 0")
	}
	return circle{centreDot: centerDot, radius: radius}, nil
}

func printArea(s Shape) {
	result := s.Area()
	fmt.Printf("Площадь фигуры: %.2f\n", result)
}

func main() {
	// 1.1
	arrIP := [4]byte{127, 0, 0, 1}
	strIP, err := formatIP(arrIP)
	if err != nil {
		fmt.Println("Ошибка: ", err)
	} else {
		fmt.Println(strIP)
	}

	// 1.1
	a := 1
	b := 10
	arrEven, err := listEven(a, b)
	if err != nil {
		fmt.Println("Ошибка: ", err)
	} else {
		fmt.Println(arrEven)
	}
	// 2
	str := "Hello"
	charCounts, err := countCharacters(str)
	if err != nil {
		fmt.Println("Ошибка: ", err)
	} else {
		fmt.Println(charCounts)
	}

	// 3
	circle, err := NewCircle(Dot{x: 0, y: 0}, 5)
	triangle, err1 := NewTriangle([3]Dot{{1, 1}, {10, 5}, {5, 10}})
	if err != nil || err1 != nil {
		fmt.Println("Ошибка: ", err, err1)
	} else {
		printArea(&triangle)
		printArea(&circle)
	}
}
