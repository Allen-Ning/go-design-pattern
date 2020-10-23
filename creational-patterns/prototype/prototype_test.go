package prototype

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestFileClone(t *testing.T) {
	file := &file{name: "file"}
	file2 := file.clone()
	assert.Equal(t, file.printName(), file2.printName())
	assert.NotEqual(t, file.printUuid(), file2.printUuid())
}

func TestFolderClone(t *testing.T) {
	file1 := &file{uuid: uuid.New().String(), name: "file1"}
	file2 := &file{uuid: uuid.New().String(), name: "file2"}
	children := []node{file1, file2}

	folder := &folder{uuid: uuid.New().String(), name: "folder", children: children}
	folder2 := folder.clone()

	assert.Equal(t, folder.printName(), folder2.printName())
	assert.NotEqual(t, folder.printUuid(), folder2.printUuid())
}
