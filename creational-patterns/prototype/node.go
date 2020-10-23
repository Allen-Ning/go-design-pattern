package prototype

type node interface {
	printName() string
	printUuid() string
	clone() node
}
