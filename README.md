# Slate To Text

Convert slatejs state to plain text

Extracted from MarkdownToSlate

# Usage

```
const paragraph = `[{type: "paragraph", children: [{text: "test paragraph"}]}]`
const text := slate_to_text.Stringify(paragraph)
// text = "test paragraph\n"
```

# License

MIT License
