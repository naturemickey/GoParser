package test_examples_go_context

import "context"

func test2(ctx context.Context, abc string) string {
	fff := func() {
		ctx.Value("ctxval")
	}
	go fff()
	return "xxxx"
}
