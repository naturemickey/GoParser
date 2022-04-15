package test_examples_go_context

import "context"

func test1(ctx context.Context, abc string) string {
	fff := func() {
		sldkafjaskdfasdf(ctx)
		// sldkafjaskdfasdf(nil)
	}
	go fff()
	return "xxxx"
}

func sldkafjaskdfasdf(ctx context.Context) {
	println(".......")
}
