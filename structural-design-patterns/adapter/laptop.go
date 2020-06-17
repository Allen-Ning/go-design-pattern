package adapter

import "fmt"

type laptop struct{}

func (m *laptop) sendEmail(info string) {
	fmt.Println("send email by laptop")
}

type laptop_adapter struct {
	laptop *laptop
}

func (adapter *laptop_adapter) send(info string) {
	adapter.laptop.sendEmail(info)
}
