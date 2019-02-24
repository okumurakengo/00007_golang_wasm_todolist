package main

import (
	"syscall/js"
)

func load() {
	document := js.Global().Get("document")
	templateElement := document.Call("getElementById", "todo").Get("content")
	
	document.
		Call("getElementById", "add").
		Call("addEventListener", "submit", js.NewEventCallback(js.PreventDefault, func(e js.Value) {
			if(document.Call("getElementById", "addText").Get("value").String() == "") {
				return
			}

			// 新しいtodo用のDOMを作成
			cloneElement := templateElement.Call("cloneNode", true)

			// チェックボックスのイベント設定
			cloneElement.
				Call("querySelector", ".todoCheck").
				Call("addEventListener", "change", js.NewEventCallback(js.PreventDefault, func(e js.Value) {
					e.Get("target").
						Get("parentNode").
						Call("querySelector", ".todoBody").
						Get("style").
						Set("text-decoration", map[bool]string{true: "line-through", false: "none"}[e.Get("target").Get("checked").Bool()])
				}))
		
			// todoの内容を設定
			cloneElement.
				Call("querySelector", ".todoBody").
				Set("textContent", document.Call("getElementById", "addText").Get("value").String())

			// 削除イベントを設定
			cloneElement.
				Call("querySelector", ".todoDelete").
				Call("addEventListener", "click", js.NewEventCallback(js.PreventDefault, func(e js.Value) {
					if(js.Global().Get("confirm").Invoke("削除しますか?").Bool()) {
						e.Get("target").
							Get("parentNode").
							Call("remove")
					}
				}))

			// 新しいtodoをhtmlに表示
			document.Call("getElementById", "todos").Call("appendChild", cloneElement)

			// 入力ボックスを空にしておく
			document.Call("getElementById", "addText").Set("value", "")
		}))
}

func main() {
	c := make(chan struct{}, 0)

	println("WASM Go Initialized")
	load()
	<-c
}
