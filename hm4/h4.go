// * Написать функцию, которая будет получать позицию коня на шахматной доске, а возвращать массив из точек, в которые конь сможет сделать ход.
// Точку следует обозначить как структуру, содержащую x и y типа int
// Получение значений и их запись в точку должны происходить только с помощью отдельных методов. В них надо проводить проверку на то, что такая точка может существовать на шахматной доске.

package main

import (
	"fmt"
	"strconv"
)

type cell struct {
	x, y int8
}

type Chessman interface {
	SetCoordinates()
	GetNextSteps() []cell
}

type Knight struct {
	cell
}

func (k *Knight) SetCoordinates() {
	var tmpInput string
	for {
		fmt.Printf("Input coordinate X(a..h): ")
		fmt.Scanln(&tmpInput)

		if len(tmpInput) != 1 {
			fmt.Println(tmpInput, len(tmpInput))
			continue
		}

		b := int8([]byte(tmpInput)[0])
		if b < int8('a') || b > int8('h') {
			continue
		}

		k.x = b - int8('a') + 1
		break
	}

	for {
		fmt.Printf("Input coordinate Y(1..8): ")
		fmt.Scanln(&tmpInput)

		coordinateY, err := strconv.ParseInt(tmpInput, 10, 64)
		if err != nil {
			continue
		}

		if coordinateY < 1 || coordinateY > 8 {
			continue
		}

		k.y = int8(coordinateY)
		break
	}
}

func (k *Knight) GetNextSteps() []cell {
	var offsets = []struct{ x, y int8 }{
		{
			x: 1, y: 2,
		},
		{
			x: 2, y: 1,
		},
		{
			x: 2, y: -1,
		},
		{
			x: 1, y: -2,
		},
		{
			x: -1, y: -2,
		},
		{
			x: -2, y: -1,
		},
		{
			x: -2, y: 1,
		},
		{
			x: -1, y: 2,
		},
	}
	nextSteps := []cell{}
	for _, offset := range offsets {
		if k.x+offset.x > 8 || k.x+offset.x < 1 ||
			k.y+offset.y > 8 || k.y+offset.y < 1 {
			continue
		}
		nextSteps = append(nextSteps, cell{x: k.x + offset.x, y: k.y + offset.y})
	}
	return nextSteps
}

func main() {
	var k Knight
	k.SetCoordinates()
	fmt.Println("Current coordinates: ", k)
	for i, cell := range k.GetNextSteps() {
		fmt.Println("Step: ", i+1, string(cell.x+int8('a')-1), cell.y)
	}
}
