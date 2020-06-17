package chain_responsibility

import "fmt"

type handlerB struct {
	nextHandler handler
}

func (h *handlerB) next(r *request, res *response) {
	fmt.Println("next2.....")
	if h.nextHandler != nil {
		h.nextHandler.process(r, res)
	}
}

func (h *handlerB) process(r *request, res *response) {
	fmt.Println("proces2.....")
	res.data["handlerB"] = "done"
	h.next(r, res)
}
