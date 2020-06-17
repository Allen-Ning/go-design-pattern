package adapter

import "fmt"

type mobile struct{}

func (m *mobile) send(info string) {
	fmt.Println("send message by mobile")
}
