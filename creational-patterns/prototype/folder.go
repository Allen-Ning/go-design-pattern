package prototype

import (
	"strings"

	"github.com/google/uuid"
)

type folder struct {
	uuid     string
	name     string
	children []node
}

func (f *folder) printName() string {
	var sb strings.Builder
	sb.WriteString("folder name: " + f.name + "|")
	for _, child := range f.children {
		sb.WriteString("file name: " + child.printName() + ",")
	}
	return sb.String()
}

func (f *folder) printUuid() string {
	return f.uuid
}

func (f *folder) clone() node {
	children := []node{}
	for _, file := range f.children {
		children = append(children, file.clone())
	}

	return &folder{uuid: uuid.New().String(), name: f.name, children: children}
}
