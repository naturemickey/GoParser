package test_examples_go_context

import . "context"

func test5(ctx1 Context, abc string) string {
	go func(ctx Context) {
		println("abcd")
	}(ctx1)
	return "xxxx"
}
