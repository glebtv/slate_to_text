package slate_to_text

type Node struct {
	Children []Node
	Type     *string
	Text     *string
	Code     *bool
}
