package visitor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVisitor(t *testing.T) {
	visitor1 := &realVisitor{base: 10}
	visitor2 := &realVisitor{base: 20}

	c := circle{radius: 5.0}
	assert.Equal(t, "circle", c.getType())
	assert.Equal(t, 785.0, c.cal(visitor1))
	assert.Equal(t, 1570.0, c.cal(visitor2))

	r := rectangle{width: 5.0, long: 6.0}
	assert.Equal(t, "rectangle", r.getType())
	assert.Equal(t, 300.0, r.cal(visitor1))
	assert.Equal(t, 600.0, r.cal(visitor2))

	s := square{size: 5.0}
	assert.Equal(t, s.getType(), "square")
	assert.Equal(t, 250.0, s.cal(visitor1))
	assert.Equal(t, 500.0, s.cal(visitor2))
}
