package main

import (
	"fmt"
	"syscall/js"

	"github.com/gofrs/uuid"
)

func main() {
	c := make(chan struct{}, 0)

	var cb js.Callback

	cb = js.NewCallback(func(args []js.Value) {
		var output string
		id, err := uuid.NewV4()
		if err != nil {
			output = fmt.Sprintf("ERR: %v", err)
		} else {
			output = id.String()
		}
		//fmt.Println(output)
		js.Global().Get("document").Call("getElementById", "myInput").Set("value", output)
	})

	js.Global().Get("document").Call("getElementById", "myButton").Call("addEventListener", "click", cb)

	<-c
}
