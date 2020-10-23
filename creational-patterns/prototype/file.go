package prototype

import "github.com/google/uuid"

type file struct {
	uuid string
	name string
}

func (f *file) printName() string {
	return f.name
}

func (f *file) printUuid() string {
	return f.uuid
}

func (f *file) clone() node {
	return &file{uuid: uuid.New().String(), name: f.name}
}
