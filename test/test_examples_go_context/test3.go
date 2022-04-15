package test_examples_go_context

import "context"

func test3(ctx context.Context, abc string) string {
	go ctx.Value("aaa")
	return "xxxx"
}
