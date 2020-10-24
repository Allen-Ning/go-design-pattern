package decorator

type MeatLover struct {
	price int
}

func (p *MeatLover) getPrice() int {
	return p.price
}
