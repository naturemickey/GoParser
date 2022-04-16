package test_examples_go_context

import "context"

type Ctx struct {
	c context.Context
}

func test8(ctx context.Context, abc string) string {
	var ctx1 = &Ctx{ctx}

	fff := func(ctx2 *Ctx) {
		println("abcd")
	}
	go fff(ctx1)
	return "xxxx"
}
