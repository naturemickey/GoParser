package main

import "GoParser/parser"

func main() {

	funcdecl := parser.ParseGoCodeToFunc("func f(ctx context.Context, a A){" +
		"a.ffff()" +
		"}")

	println(funcdecl.String())

}
