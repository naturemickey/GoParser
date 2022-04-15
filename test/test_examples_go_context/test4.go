package test_examples_go_context

import "context"

func test4(ctx context.Context, abc string) string {
	fff := func(ctx context.Context) {
		println("abcd")
	}
	go fff(ctx)
	return "xxxx"
}
