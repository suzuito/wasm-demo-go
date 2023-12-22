package main

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"runtime"
	"syscall/js"
)

func main() {
	// 注意 サンプルコードなので引数のエラーチェックは省略してます。

	// Goの関数をJavaScript中で呼ぶサンプルコード
	// Set関数を使う。この関数は、JavaScriptオブジェクトのプロパティを追加する。
	js.Global().Set("hoge", js.ValueOf("This value is set in Go."))
	// Globalオブジェクト(window, global)のmd5sumプロパティにGoの関数をセットする。
	js.Global().Set(
		"md5sum",
		// 入力された文字列のmd5sumを計算する関数
		js.FuncOf(func(this js.Value, args []js.Value) any {
			input := args[0].String()
			if len(input) <= 0 {
				// 空文字列が入力された場合はエラーを返す
				return map[string]any{
					"error": "ERROR empty input",
				}
			}
			hash := md5.Sum([]byte(input))
			return map[string]any{
				"md5": hex.EncodeToString(hash[:]),
			}
		}),
	)

	// JavaScriptの関数やプロパティをGoから呼ぶサンプルコード
	// Getメソッドがjsオブジェクトのプロパティを取得し、Callがjsオブジェクトのメソッドを呼ぶ。
	{
		// Goのランタイム情報をHTMLへ出力する
		output := fmt.Sprintf("Runtime version=%s", runtime.Version())
		js.Global().Get("document").Call("getElementById", "go-info").Set("innerHTML", output)
	}
	{
		// クリックイベントをbindする
		js.Global().Get("document").Call("getElementById", "sha256-button").
			Call("addEventListener", "click", js.FuncOf(func(this js.Value, args []js.Value) any {
				input := js.Global().Get("document").Call("getElementById", "sha256-input").Get("value").String()
				if len(input) <= 0 {
					// 空文字列が入力された場合はエラーを返す
					js.Global().Get("document").Call("getElementById", "sha256-error").Set("innerHTML", "ERROR empty input")
					return nil
				}
				hash := sha256.New().Sum([]byte(input))
				js.Global().Get("document").Call("getElementById", "sha256-output").Set("value", hex.EncodeToString(hash[:]))
				return nil
			}))
	}

	// main関数が終了すると、Instanceがなくなってしまう。
	// 以下はmain関数の終了を防止するためのブロック。
	blocker := make(chan struct{})
	<-blocker
}
