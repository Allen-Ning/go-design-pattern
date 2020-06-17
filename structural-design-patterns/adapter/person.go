package adapter

type sender interface {
	send(string)
}

type person struct{}

func (p *person) communicate_through(s sender, info string) {
	s.send(info)
}
