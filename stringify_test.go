package slate_to_text

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func RunStringify(t *testing.T, input string) string {
	s, err := Stringify([]byte(input))
	if err != nil {
		t.Fatal(err)
	}
	return string(s)
}

func TestSimple(t *testing.T) {
	s := RunStringify(t, `[{"type": "paragraph", "children": [{"text": "test"}]}]`)
	assert.Equal(t, "test\n", s, "")
}

const multiExample = `[{"type":"paragraph","children":[{"bold":true,"text":"fasd"},{"text":" "},{"text":"fasdf","italic":true},{"text":" "},{"text":"asfas","underline":true},{"text":" "},{"text":"ddfsa","strikethrough":true},{"text":" fasd fsa"}]},{"type":"heading-one","children":[{"text":"test"}]},{"type":"heading-two","children":[{"text":"test 2"}]},{"type":"heading-three","children":[{"text":"test 3 "}]},{"type":"heading-three","children":[{"code":true,"text":"inline"}]},{"type":"code","children":[{"text":"block code"}]},{"type":"paragraph","children":[{"text":""}]},{"type":"paragraph","children":[{"text":""}]}]`
const multiResult = `fasd fasdf asfas ddfsa fasd fsa
test
test 2
test 3
inline
block code
block code
`

func TesMulti(t *testing.T) {
	s := RunStringify(t, multiExample)
	assert.Equal(t, multiResult, s, "")
}
