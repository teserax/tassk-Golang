//Напишите метод String для типа coordinate
// и location и используйте его для отображения
//координат в более читабельном формате.
package main

import "fmt"

type coordinate struct {
	d, m, s float64
	h       rune
}

type stringer interface {
	String() string
}

// location с широтой и долготой в десятичных градусах
type location struct {
	lat, long float64
}

// String форматирует location с широтой и долготой
func (l location) String() string {
	return fmt.Sprintf("%v, %v", l.lat, l.long)
}
func (l location) GetString() string {

	return fmt.Sprintf("%v, %v", l.lat, l.long)
}

func (c coordinate) String() string {
	return fmt.Sprintf("%v,%v,%v,%v", c.d, c.m, c.s, c.h)
}
func info(s stringer) {
	fmt.Println(s.String())
}
func main() {
	curiosity := location{-4.5895, 137.4417}
	curiosityCoor := coordinate{4.58, 137.4, 56, 'H'}
	info(curiosity)
	info(curiosityCoor)
	fmt.Println(curiosity.GetString())
}
