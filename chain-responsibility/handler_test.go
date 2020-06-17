package chain_responsibility

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProces(t *testing.T) {
	var b *handlerB = &handlerB{nextHandler: nil}
	var a *handlerA = &handlerA{nextHandler: b}
	var r *request = &request{id: 1234}
	var res *response = &response{id: 222, data: make(map[string]string)}

	a.process(r, res)
	assert.Equal(t, res.data["handlerA"], "done", "The two words should be the same.")
	assert.Equal(t, res.data["handlerB"], "done", "The two words should be the same.")
}
