package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	fmt.Println("======Read======")
	Read()
	fmt.Println("======ReadByte======")
	ReadByte()
	fmt.Println("======ReadRune======")
	ReadRune()
	fmt.Println("======ReadBytes======")
	ReadBytes()
	fmt.Println("======ReadFrom======")
	ReadFrom()
	fmt.Println("======Reset======")
	Reset()
}

func Read() {
	bufs := bytes.NewBufferString("Learning swift.")
	fmt.Println(bufs.String())

	// 声明一个空的slice，长度和容量为8
	l := make([]byte, 8)
	// 把bufs的内容读入到l内，因为l容量为8,所以只能读8个字节过来
	bufs.Read(l)
	fmt.Println("::bufs缓冲器内容::")
	fmt.Println(bufs.String())
	// 空的l被写入了8个字符，所以为Learning
	fmt.Println("::l的slice内容::")
	fmt.Println(string(l))

	// 把bufs的内容读入到l内，原来的l的内容被覆盖了
	bufs.Read(l)
	fmt.Println("::bufs缓冲器被第二次读取后剩余的内容::")
	fmt.Println(bufs.String())
	fmt.Println("::l的slice内容被覆盖，由于bufs只有7个了，因此最后一个g被留下来了::")
	fmt.Println(string(l))
}

func ReadByte() {
	bufs := bytes.NewBufferString("Learning swift.")
	fmt.Println(bufs.String())

	c, _ := bufs.ReadByte()
	fmt.Println(string(c))
	fmt.Println(bufs.String())
}

func ReadRune() {
	bufs := bytes.NewBufferString("学swift.")
	fmt.Println(bufs.String())

	// 读取第一个rune，赋值给r
	r, z, _ := bufs.ReadRune()
	fmt.Println(bufs.String())
	fmt.Printf("r=%s, z=%d\n", string(r), z)
}

func ReadBytes() {
	bufs := bytes.NewBufferString("现在开始 Learning swiftL.")
	fmt.Println(bufs.String())

	var delim byte = 'L'
	line, _ := bufs.ReadBytes(delim)
	fmt.Println(bufs.String())
	fmt.Println(string(line))

	line, _ = bufs.ReadBytes(delim)
	fmt.Println(bufs.String())
	fmt.Println(string(line))
}

func ReadFrom() {
	file, _ := os.Open("text.txt")
	buf := bytes.NewBufferString("Learning swift.")
	buf.ReadFrom(file) //将text.txt内容追加到缓冲器的尾部
	fmt.Println(buf.String())
}

func Reset() {
	bufs := bytes.NewBufferString("现在开始 Learning swift.")
	fmt.Println(bufs.String())

	bufs.Reset()
	fmt.Println("::已经清空了bufs的缓冲内容::")
	fmt.Println(bufs.String())
}
