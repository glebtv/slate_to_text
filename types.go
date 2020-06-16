package slate_to_text

type Node struct {
	Children      []Node
	Type          *string
	Text          *string
	Code          *bool
	Italic        *bool
	Bold          *bool
	Underline     *bool
	Strikethrough *bool
	Data          map[string]interface{} `json:"data,omitempty"`
}
