package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func main() {
	// newBuffer整型转换成字节
	var n int = 10000
	intToBytes := IntToBytes(n)
	fmt.Println("==========int to byte===========")
	fmt.Println(intToBytes)
	TestBufferString()
	BufferWrite()
	BufferWriteString()
}

func IntToBytes(n int) []byte {
	x := int32(n)
	// 创建一个内容是[]byte的slice的缓冲器
	// 与bytes.NewBufferString("")等效
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

func TestBufferString() {
	buf1 := bytes.NewBufferString("swift")
	buf2 := bytes.NewBuffer([]byte("swift"))
	buf3 := bytes.NewBuffer([]byte{'s', 'w', 'i', 'f', 't'})
	fmt.Println("===============以下buf1,buf2,buf3等效==============")
	fmt.Println("buf1:", buf1)
	fmt.Println("buf2:", buf2)
	fmt.Println("buf3:", buf3)
	fmt.Println("===========以下创建空的缓冲器等效=========")
	buf4 := bytes.NewBufferString("")
	buf5 := bytes.NewBuffer([]byte{})
	fmt.Println("buf4:", buf4)
	fmt.Println("buf5:", buf5)
}

func BufferWrite() {
	fmt.Println("===========以下通过Write把swift写入Learning缓冲器尾部=========")
	bufs := bytes.NewBufferString("Learning")
	newBytes := []byte("swift")

	fmt.Println(bufs.String())
	bufs.Write(newBytes)
	fmt.Println(bufs.String())
}

func BufferWriteString() {
	bufs := bytes.NewBufferString("Learning")
	newStr := "swift"

	fmt.Println(bufs.String())
	bufs.WriteString(newStr)
	fmt.Println(bufs.String())
}
