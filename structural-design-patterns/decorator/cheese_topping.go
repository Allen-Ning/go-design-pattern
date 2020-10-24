package decorator

type cheeseTopping struct {
	pizza pizza
	price int
}

func (t *cheeseTopping) getPrice() int {
	return t.pizza.getPrice() + t.price
}
