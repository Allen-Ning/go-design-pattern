package chain_responsibility

type request struct {
	id int
}

type response struct {
	id   int
	data map[string]string
}

type handler interface {
	next(*request, *response)
	process(*request, *response)
}
