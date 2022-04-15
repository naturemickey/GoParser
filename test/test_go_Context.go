package main

import "GoParser/test/semantic"

func main() {
	semantic.WaklTest("./test/test_examples_go_context/test1.go")
	semantic.WaklTest("./test/test_examples_go_context/test2.go")
	semantic.WaklTest("./test/test_examples_go_context/test3.go")
	semantic.WaklTest("./test/test_examples_go_context/test4.go")
	semantic.WaklTest("/Users/mickey/git/mis-backend/mis_admin/src/biz/review/review_biz.go")
	semantic.WaklTest("/Users/mickey/git/mis-backend/mis_worker/src/service/mis_worker.go")
}
