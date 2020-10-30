package visitor

type visitor interface {
	calCircle(*circle) float64
	calRectangle(*rectangle) float64
	calSquare(*square) float64
}

type realVisitor struct {
	base float64
}

func (v *realVisitor) calCircle(c *circle) float64 {
	const pi float64 = 3.14
	return v.base * c.radius * c.radius * pi
}

func (v *realVisitor) calRectangle(r *rectangle) float64 {
	return v.base * r.long * r.width
}

func (v *realVisitor) calSquare(s *square) float64 {
	return v.base * s.size * s.size
}
