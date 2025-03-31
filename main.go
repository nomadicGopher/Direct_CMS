package main

import (
	"fmt"
)

func main() {
	fmt.Println("This returns OK in the DOM.")
	fmt.Println("This does NOT return OK in the DOM.") //!
	select {}                                          // Keep go running in the background via WASM.
}
