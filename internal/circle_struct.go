package internal

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	radius := c.Radius
	if radius < 0 {
		radius = -radius // Make radius positive
	}
	return 3.14159 * radius * radius
}

func (c Circle) Perimeter() float64 {
	radius := c.Radius
	if radius < 0 {
		radius = -radius // Make radius positive
	}
	return 2 * 3.14159 * radius
}
