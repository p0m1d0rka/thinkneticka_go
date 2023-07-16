package hw

import (
	"math"
)

// По условиям задачи, координаты не могут быть меньше 0.

// Добавляем структуру точки
type Point struct {
	X float64
	Y float64
}

// меняем геом, теперь он наследует от точки
type Geom struct {
	p1 Point
	p2 Point
}

func (geom Geom) CalculateDistance() (distance float64) {
	// так как по условию задачи координаты строго больше нуля, то эту проверку можно опустить
	distance = math.Sqrt(math.Pow(geom.p2.X-geom.p1.X, 2) + math.Pow(geom.p2.Y-geom.p1.Y, 2))

	// возврат расстояния между точками
	return distance
}
