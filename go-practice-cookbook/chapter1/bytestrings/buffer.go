package bytestrings

import (
	"bytes"
	"io"
	"io/ioutil"
)

// Buffer demostrates some tricks for initializing bytes
// Buffers
// This buffers implement an io,Reader interface
func Buffer(rawString string) *bytes.Buffer {
	//we will start with a string encoded into raw bytes
	rawBytes := []byte(rawString)
	//there are number of ways to create a buffer from
	//the raw bytes or from original stiring
	var b = new(bytes.Buffer)
	b.Write(rawBytes)

	// alternatively
	//b = bytes.NewBuffer(rawBytes)
	// and avoiding the initial byte array altogether
	b = bytes.NewBufferString(rawString)
	return b
}

// TpString is an example of yalking an io,Reader and consuming
// it all, then returning a string

func toString(r io.Reader) (string, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
