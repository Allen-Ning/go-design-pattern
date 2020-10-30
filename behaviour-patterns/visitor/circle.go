package visitor

type shape interface {
	cal(*visitor) int
	getType() string
}

type circle struct {
	radius float64
}

func (c *circle) cal(v visitor) float64 {
	return v.calCircle(c)
}

func (c *circle) getType() string {
	return "circle"
}

type rectangle struct {
	width float64
	long  float64
}

func (r *rectangle) getType() string {
	return "rectangle"
}

func (r *rectangle) cal(v visitor) float64 {
	return v.calRectangle(r)
}

type square struct {
	size float64
}

func (s *square) getType() string {
	return "square"
}

func (s *square) cal(v visitor) float64 {
	return v.calSquare(s)
}
