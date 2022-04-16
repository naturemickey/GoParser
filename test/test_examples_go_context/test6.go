package test_examples_go_context

import "context"

type CtxCtx context.Context

func test6(ctx CtxCtx, abc string) string {
	fff := func() {
		ctx.Value("ctxval")
	}
	go fff()
	return "xxxx"
}
