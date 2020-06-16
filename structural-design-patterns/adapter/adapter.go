package adapter

import "fmt"

type laptop struct {
}

func (m *laptop) sendEmail(info string) {
	fmt.Println("send email by laptop")
}

type laptop_adapter struct {
	laptop *laptop
}

func (adapter *laptop_adapter) send(info string) {
	adapter.laptop.sendEmail(info)
}

type post_office_adapter struct {
	post_office *post_office
}

func (adapter *post_office_adapter) send(info string) {
	adapter.post_office.post(info)
}

type post_office struct{}

func (m *post_office) post(info string) {
	fmt.Println("send mail by post office")
}

type mobile struct{}

func (m *mobile) send(info string) {
	fmt.Println("send message by mobile")
}

type sender interface {
	send(string)
}

type person struct{}

func (p *person) communicate_through(s sender, info string) {
	s.send(info)
}
