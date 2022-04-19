package test_examples_go_context

import "context"

type Ctx struct {
	c context.Context
}

func test8(ctx context.Context, abc string) string {
	var ctx1 = &Ctx{ctx}

	fff := func() {
		ctx1.c.Value("ctxval")
	}
	go fff()
	return "xxxx"
}
