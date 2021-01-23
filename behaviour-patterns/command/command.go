package command

type Command interface {
	do(Content)
	undo(Content)
}

type Content struct {
	text string
}

type AddPrefixCommand struct {
	prefix string
}

func (c *AddPrefixCommand) do(content *Content) {
	content.text = c.prefix + content.text
}

func (c *AddPrefixCommand) undo(content *Content) {
	content.text = content.text[len(c.prefix):]
}
