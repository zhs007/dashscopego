package qwen

// TextConent is used for text-generation only.
type TextContent struct {
	Text string
}

var _ IQwenContentMethods = &TextContent{}

func NewTextContent() *TextContent {
	t := TextContent{""}
	return &t
}

func (t *TextContent) ToBytes() []byte {
	return []byte(t.Text)
}

func (t *TextContent) ToString() string {
	return t.Text
}

func (t *TextContent) SetText(text string) {
	if t == nil {
		panic("TextContent is nil")
	}
	t.Text = text
}

func (t *TextContent) AppendText(text string) {
	if t == nil {
		panic("TextContent is nil")
	}
	t.Text += text
}

func (t *TextContent) SetImage(_ string) {
	panic("TextContent does not support SetImage")
}

func (t *TextContent) SetAudio(_ string) {
	panic("TextContent does not support SetAudio")
}

// func foo() {
// 	a := &TextContent{}
// 	fmt.Println(a)
// }
