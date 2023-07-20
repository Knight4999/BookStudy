package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	var w io.Writer
	w = os.Stdout
	f := w.(*os.File)
	fmt.Printf("f: %T,%v\n", f, f)
	/* c := w.(*bytes.Buffer)
	fmt.Printf("c: %v\n", c) */

	rw := w.(io.ReadWriter)
	fmt.Printf("rw: %T,%v\n", rw, rw)

}
