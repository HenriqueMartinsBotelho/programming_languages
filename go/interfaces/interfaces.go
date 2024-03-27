package interfaces

import (
	"fmt"
	"math"
)

/*
Você pode dizer que rect e circle são do tipo geometry. Isso é possível graças à forma como as interfaces são implementadas em Go.
Uma interface, neste caso geometry, é um tipo que especifica um conjunto de métodos,
mas não implementa esses métodos por si só.
Um tipo é considerado como implementando uma interface apenas se ele possui todos os métodos que a interface declara.
*/

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func TestCode() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)
}
