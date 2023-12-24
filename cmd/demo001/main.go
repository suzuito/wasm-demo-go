package main

import (
	"fmt"
	"runtime"
	"syscall/js"
)

func main() {
	// ブラウザ上で実行される処理は、必ずmain関数を介さなければならない。
	// main関数を介さずに、ある特定の関数だけをブラウザ上で実行する、みたいなことはできないらしい(この辺、RustやC＋＋だとどうなんだろう？知ってる人いたら教えて下さい)。

	// 標準出力は、デベロッパーツールのConsoleへ出力される
	fmt.Println("Hello, WebAssembly!!!")

	// 標準ライブラリsyscall/jsを使って、ブラウザ上のJavaScriptオブジェクトを操作する。
	output := fmt.Sprintf("Runtime version=%s", runtime.Version())
	js.Global().Get("document").Call("getElementById", "go-info").Set("innerHTML", output)

	// main関数が終了すると、Instanceがなくなってしまう。
	// 以下はmain関数の終了を防止するためのブロック。
	blocker := make(chan struct{})
	<-blocker
}
