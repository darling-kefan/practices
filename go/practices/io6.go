package main

import (
	_ "bufio"
	"bytes"
	"fmt"
	"io"
)

// data represents a table of input and expected output.
var data = []struct {
	input  []byte
	output []byte
}{
	// {[]byte("abc"), []byte("abc")},
	// {[]byte("elvis"), []byte("Elvis")},
	// {[]byte("aElvis"), []byte("aElvis")},
	{[]byte("abcelvis"), []byte("abcElvis")},
	{[]byte("eelvis"), []byte("eElvis")},
	{[]byte("aelvis"), []byte("aElvis")},
	{[]byte("aabeeeelvis"), []byte("aabeeeElvis")},
	{[]byte("e l v i s"), []byte("e l v i s")},
	{[]byte("aa bb e l v i saa"), []byte("aa bb e l v i saa")},
	{[]byte(" elvi s"), []byte(" elvi s")},
	{[]byte("elvielvis"), []byte("elviElvis")},
	{[]byte("elvielvielviselvi1"), []byte("elvielviElviselvi1")},
	{[]byte("elvielviselvis"), []byte("elviElvisElvis")},
}

// Declare what needs to be found and its replacement.
var find = []byte("elvis")
var repl = []byte("Elvis")

// Calculate the number of bytes we need to locate.
var size = len(find)

func main() {
	var output bytes.Buffer

	fmt.Println("=======================================\nRunning Algorithm One")
	for _, d := range data {
		input := bytes.NewReader(d.input)
		output.Reset()
		algOne(input, &output)

		matched := bytes.Compare(d.output, output.Bytes())
		fmt.Printf("Matched: %v Inp: [%s] Exp: [%s] Got: [%s]\n",
			matched == 0,
			d.input,
			d.output,
			output.Bytes())

		break
	}
}

// algOne is one way to solve the problem. This approach first
// reads a mininum number of bytes required and then starts processing
// new bytes as they are provided in the stream.
func algOne(r io.Reader, w *bytes.Buffer) {
	// Declare the buffers we need to process the stream.
	buf := make([]byte, size)
	tmp := make([]byte, 1)
	end := size - 1

	// Read in an initial number of bytes we need to get started.
	if n, err := io.ReadFull(r, buf[:end]); err != nil {
		w.Write(buf[:n])
		return
	}

	for {
		// Read in one byte from the input stream.
		n, err := io.ReadFull(r, tmp)

		// If we have a byte then process it.
		if n == 1 {
			// Add this byte to the end of the buffer.
			buf[end] = tmp[0]

			fmt.Println(buf, end)

			// If we have a match, replace the bytes.
			if bytes.Compare(buf, find) == 0 {
				copy(buf, repl)
			}

			// Write the front byte since it has been compared.
			w.WriteByte(buf[0])

			// Slice that front byte out.
			copy(buf, buf[1:])
		}

		// Did we hit the end of the stream, then we are done.
		if err != nil {
			// Flush the reset of the bytes we have.
			w.Write(buf[:end])
			break
		}
	}
}
