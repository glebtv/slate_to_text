package slate_to_text

import (
	"bytes"
	"encoding/json"
)

func StringifyNodes(buf *bytes.Buffer, nodes []Node) {
	for _, node := range nodes {
		if node.Text != nil {
			buf.WriteString(*node.Text + "\n")
		}
		if node.Children != nil {
			StringifyNodes(buf, node.Children)
		}
	}
}

func Stringify(input []byte) (string, error) {
	data := make([]Node, 0)
	err := json.Unmarshal(input, &data)
	if err != nil {
		return "", err
	}

	var ret bytes.Buffer
	StringifyNodes(&ret, data)
	return ret.String(), nil
}
