package decorator

type tomatoTopping struct {
	pizza pizza
	price int
}

func (t *tomatoTopping) getPrice() int {
	return t.pizza.getPrice() + t.price
}
