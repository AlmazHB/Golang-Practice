package bytestrings

import (
	"bufio"
	"bytes"
	"fmt"
)

// Working with Buffer will make use of
// the buffer created by the Buffer function
func WorkWithBuffer() error {
	rawString := "it's easy to encode unicode into a byte array"
	b := Buffer(rawString)
	// we can quickly convert a buffer back into byes with
	//b.Bytes() or a string with b.String()
	fmt.Println(b.String())
	//because this is an io Reader we can make use of
	// generic io readre functions such as
	s, err := toString(b)
	if err != nil {
		return err
	}
	fmt.Println(s)
	// we can also take our bytes and create a byte reader
	//these readers implement io.Reader, io.ReaderAt,
	//io.WriteTo, io.Seeker, io.Byte Scanner and
	//io.RuneScanner interface
	reader := bytes.NewReader([]byte(rawString))

	//we can also plug it into a scanner
	//buffered reading and tokenzation

	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Print(scanner.Text())
	}
	return nil
}
