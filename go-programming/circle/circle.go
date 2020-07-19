package circle

type Circle struct {
	radius float64
}

func (c Circle) GetArea() float64 {
	return 3.14 * c.radius * c.radius
}
