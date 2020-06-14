package main

import "fmt"

type handlerA struct {
	nextHandler handler
}

func (h *handlerA) next(r *request, res *response) {
	fmt.Println("next1.....")
	if h.nextHandler != nil {
		h.nextHandler.process(r, res)
	}
}

func (h *handlerA) process(r *request, res *response) {
	fmt.Println("proces1.....")
	res.data["handlerA"] = "done"
	h.next(r, res)
}
