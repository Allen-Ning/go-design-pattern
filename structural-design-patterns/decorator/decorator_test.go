package decorator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecorator(t *testing.T) {
	pizza := &MeatLover{price: 15}
	pizza_with_cheese := &cheeseTopping{pizza: pizza, price: 2}
	pizza_with_cheese_and_tomato := &tomatoTopping{pizza: pizza_with_cheese, price: 3}

	assert.Equal(t, pizza_with_cheese_and_tomato.getPrice(), 20)
}
