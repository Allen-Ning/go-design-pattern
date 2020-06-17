package adapter

import "fmt"

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
