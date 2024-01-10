package life

import (
	"fmt"
	"math/rand"
	"time"
)

type World struct {
	Height int // Высота сетки
	Width  int // Ширина сетки
	Cells  [][]bool
}

// Используйте код из предыдущего урока по игре «Жизнь»
func NewWorld(height, width int) (*World, error) {
	// создаём тип World с количеством слайсов hight (количество строк)
	if height <= 0 || width <= 0 {
		return nil, fmt.Errorf("New err")
	}
	cells := make([][]bool, height)
	for i := range cells {
		cells[i] = make([]bool, width) // создаём новый слайс в каждой строке
	}
	return &World{
		Height: height,
		Width:  width,
		Cells:  cells,
	}, nil
}
func (w *World) next(x, y int) bool {
	n := w.neighbors(x, y)       // получим количество живых соседей
	alive := w.Cells[y][x]       // текущее состояние клетки
	if n < 4 && n > 1 && alive { // если соседей двое или трое, а клетка жива
		return true // то следующее состояние — жива
	}
	if n == 3 && !alive { // если клетка мертва, но у неё трое соседей
		return true // клетка оживает
	}

	return false // в любых других случаях — клетка мертва
}
func (w *World) neighbors(x, y int) int {
	f := func(x bool) int {
		if x == true {
			return 1
		} else {
			return 0
		}
	}
	nei := 0
	if y == 0 && x == 0 {
		nei = f(w.Cells[0][1]) + f(w.Cells[1][1]) + f(w.Cells[1][0]) + f(w.Cells[0][w.Width-1]) + f(w.Cells[1][w.Width-1]) +
			f(w.Cells[w.Height-1][0]) + f(w.Cells[w.Height-1][1]) + f(w.Cells[w.Height-1][w.Width-1])
	} else if y == w.Height-1 && x == 0 {
		nei = f(w.Cells[y-1][0]) + f(w.Cells[y-1][1]) + f(w.Cells[y][1]) + f(w.Cells[0][0]) + f(w.Cells[0][1]) +
			f(w.Cells[0][w.Width-1]) + f(w.Cells[w.Height-1][w.Width-1]) + f(w.Cells[w.Height-2][w.Width-1])
	} else if y == 0 && x == w.Width-1 {
		nei = f(w.Cells[0][x-1]) + f(w.Cells[1][x-1]) + f(w.Cells[1][x]) + f(w.Cells[0][0]) + f(w.Cells[1][0]) +
			f(w.Cells[w.Height-1][w.Width-1]) + f(w.Cells[w.Height-1][w.Width-2]) + f(w.Cells[w.Height-1][0])
	} else if y == w.Height-1 && x == w.Width-1 {
		nei = f(w.Cells[y][x-1]) + f(w.Cells[y-1][x-1]) + f(w.Cells[y-1][x]) + f(w.Cells[0][w.Width-1]) +
			f(w.Cells[0][w.Width-2]) + f(w.Cells[w.Height-1][0]) + f(w.Cells[w.Height-2][0]) + f(w.Cells[0][0])
	} else if y == 0 {
		nei = f(w.Cells[0][x-1]) + f(w.Cells[0][x+1]) + f(w.Cells[1][x-1]) + f(w.Cells[1][x]) + f(w.Cells[1][x+1]) +
			f(w.Cells[w.Height-1][x-1]) + f(w.Cells[w.Height-1][x]) + f(w.Cells[w.Height-1][x+1])
	} else if x == 0 {
		nei = f(w.Cells[y-1][0]) + f(w.Cells[y-1][1]) + f(w.Cells[y][1]) + f(w.Cells[y+1][1]) + f(w.Cells[y+1][0]) +
			f(w.Cells[y-1][w.Width-1]) + f(w.Cells[y][w.Width-1]) + f(w.Cells[y+1][w.Width-1])
	} else if y == w.Height-1 {
		nei = f(w.Cells[y][x-1]) + f(w.Cells[y-1][x-1]) + f(w.Cells[y-1][x]) + f(w.Cells[y-1][x+1]) + f(w.Cells[y][x+1]) +
			f(w.Cells[0][x-1]) + f(w.Cells[0][x]) + f(w.Cells[0][x+1])
	} else if x == w.Width-1 {
		nei = f(w.Cells[y-1][x]) + f(w.Cells[y-1][x-1]) + f(w.Cells[y][x-1]) + f(w.Cells[y+1][x-1]) + f(w.Cells[y+1][x]) +
			f(w.Cells[y-1][0]) + f(w.Cells[y][0]) + f(w.Cells[y+1][0])
	} else {
		nei = f(w.Cells[y-1][x-1]) + f(w.Cells[y-1][x]) + f(w.Cells[y-1][x+1]) + f(w.Cells[y][x-1]) +
			f(w.Cells[y][x]) + f(w.Cells[y][x+1]) + f(w.Cells[y+1][x-1]) + f(w.Cells[y+1][x]) + f(w.Cells[y+1][x+1])
	}
	return nei
}
func NextState(oldWorld, newWorld *World) {
	// переберём все клетки, чтобы понять, в каком они состоянии
	for i := 0; i < oldWorld.Height; i++ {
		for j := 0; j < oldWorld.Width; j++ {
			// для каждой клетки получим новое состояние
			newWorld.Cells[i][j] = oldWorld.next(j, i)
		}
	}
}

// RandInit заполняет поля на указанное число процентов
func (w *World) RandInit(percentage int) {
	// Количество живых клеток
	numAlive := percentage * w.Height * w.Width / 100
	// Заполним живыми первые клетки
	w.fillAlive(numAlive)
	// Получаем рандомные числа
	r := rand.New(rand.NewSource(time.Now().Unix()))

	// Рандомно меняем местами
	for i := 0; i < w.Height*w.Width; i++ {
		randRowLeft := r.Intn(w.Width)
		randColLeft := r.Intn(w.Height)
		randRowRight := r.Intn(w.Width)
		randColRight := r.Intn(w.Height)

		w.Cells[randRowLeft][randColLeft] = w.Cells[randRowRight][randColRight]
	}
}

func (w *World) fillAlive(num int) {
	aliveCount := 0
	for j, row := range w.Cells {
		for k := range row {
			w.Cells[j][k] = true
			aliveCount++
			if aliveCount == num {

				return
			}
		}
	}
}
