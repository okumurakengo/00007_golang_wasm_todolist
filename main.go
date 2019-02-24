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

			cloneElement := templateElement.Call("cloneNode", true)

			cloneElement.
				Call("querySelector", ".todoCheck").
				Call("addEventListener", "change", js.NewEventCallback(js.PreventDefault, func(e js.Value) {
					e.Get("target").
						Get("parentNode").
						Call("querySelector", ".todoBody").
						Get("style").
						Set("text-decoration", map[bool]string{true: "line-through", false: "none"}[e.Get("target").Get("checked").Bool()])
				}))
		
			cloneElement.
				Call("querySelector", ".todoBody").
				Set("textContent", document.Call("getElementById", "addText").Get("value").String())

			cloneElement.
				Call("querySelector", ".todoDelete").
				Call("addEventListener", "click", js.NewEventCallback(js.PreventDefault, func(e js.Value) {
					if(js.Global().Get("confirm").Invoke("削除しますか?").Bool()) {
						e.Get("target").
							Get("parentNode").
							Call("remove")
					}
				}))

			document.Call("getElementById", "todos").Call("appendChild", cloneElement)

			document.Call("getElementById", "addText").Set("value", "")
		}))
}

func main() {
	c := make(chan struct{}, 0)

	println("WASM Go Initialized")
	load()
	<-c
}
