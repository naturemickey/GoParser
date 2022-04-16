package test_examples_go_context

import "context"

func test7(ctx context.Context, abc string) string {
	type Ctxccc context.Context

	var ctx1 Ctxccc = ctx

	fff := func(ctx2 Ctxccc) {
		println("abcd")
	}
	go fff(ctx1)
	return "xxxx"
}
