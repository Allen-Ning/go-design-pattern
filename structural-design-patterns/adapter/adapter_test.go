package adapter

import "testing"

func test(t *testing.T) {
	p := &person{}
	mobile := &mobile{}
	laptop := &laptop_adapter{}
	post_office := &post_office_adapter{}
	p.communicate_through(mobile, "haha")
	p.communicate_through(laptop, "haha")
	p.communicate_through(post_office, "haha")
}
