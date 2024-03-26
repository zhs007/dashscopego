package main

import (
	"context"
	"log"
	"os"

	"github.com/zhs007/dashscopego"
	"github.com/zhs007/dashscopego/qwen"
)

func main() {
	model := qwen.QwenTurbo
	token := os.Getenv("DASHSCOPE_API_KEY")

	if token == "" {
		panic("token is empty")
	}

	cli := dashscopego.NewTongyiClient(model, token)

	// content := qwen.TextContent{Text: `[{"text": "告诉我番茄炒蛋的做法"}]`}
	content := qwen.TextContent{
		Text:  `[{"text": "总结文档"},{"file": "https://qianwen-res.oss-cn-beijing.aliyuncs.com/QWEN_TECHNICAL_REPORT.pdf"}]`,
		IsRaw: true,
	}

	input := dashscopego.TextInput{
		Messages: []dashscopego.TextMessage{
			{Role: "user", Content: &content},
		},
	}

	// callback function:  print stream result
	streamCallbackFn := func(ctx context.Context, chunk []byte) error {
		log.Print(string(chunk))
		return nil
	}
	req := &dashscopego.TextRequest{
		Input:       input,
		StreamingFn: streamCallbackFn,
		Plugin:      `{"pdf_extracter":{}}`,
		// HasUploadOss: true,
	}

	ctx := context.TODO()
	resp, err := cli.CreateCompletion(ctx, req)
	if err != nil {
		panic(err)
	}

	log.Println("\nnon-stream result: ")
	log.Println(resp.Output.Choices[0].Message.Content.ToString())
}
