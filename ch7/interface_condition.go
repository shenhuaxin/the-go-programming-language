package main

import (
	"bytes"
	"io"
	"os"
)

func main() {

	var w io.Writer

	w = os.Stdout
	w = new(bytes.Buffer)
	//w = time.Second
	w.Write([]byte("hello"))
	var rwc io.ReadWriteCloser
	rwc = os.Stdout
	rwc.Write([]byte("hello"))
	//rwc = new(bytes.Buffer)
}
