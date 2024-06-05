package interfaces

import (
	"io"
	"os"
)

// PipeExample helps give some more examl=ples of using io
//interfaces

func PipeExample() error {
	// the pipe reader and pipe writer implement
	// io.Reader and io.Writer
	r, w := io.Pipe()
	//this needs to be run in a separate go routine
	//as will block waiting for the reader
	//close at the end for cleanup
	go func() {
		// for now we will write something basic
		//this cloud also be used to encode json
		//base64 encode, etc.
		w.Write([]byte("test\n"))
		w.Close()

	}()
	if _, err := io.Copy(os.Stdout, r); err != nil {
		return err
	}
	return nil
}
