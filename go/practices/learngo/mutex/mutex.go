// test for Go
//
// Copyright (c) 2015 - Batu
//
// 创建一个文件存放数据，在同一时刻，可能会有多个Goroutine分别对此文件进行写操作和读操作
// 每一次写操作都应该向这个文件写入若干的数据，作为一个独立的数据块存在，这意味着写操作之间不能彼此干扰，写入的内容之间也不能出现穿插和混淆的情况
// 每一次读操作都应该从这个文件中读取一个独立完整的数据块。它们读取的数据块不能重复，且需要按顺序读取。
// 例如：第一个读操作读取了数据块1,第二个操作就应该读取数据块2,第三个读操作则应该读取数据块3,以此类推
// 对于这些读操作是否可以被同时执行，不做要求。即使同时进行，也应该保持先后顺序。
package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

// 数据文件的接口类型
type DataFile interface {
	// 读取一个数据块
	Read(complete <-chan bool) (rsn int64, d Data, err error)
	// 写入一个数据块
	Write(d Data) (wsn int64, err error)
	// 获取最后读取的数据块的序列号
	Rsn() int64
	// 获取最后写入的数据块的序列号
	Wsn() int64
	// 获取数据块的长度
	DataLen() uint32
}

// 数据类型
type Data []byte

// 数据文件的实现类型
type myDataFile struct {
	f       *os.File     //文件
	fmutex  sync.RWMutex //被用于文件的读写锁
	woffset int64        //写操作需要用到的偏移量
	roffset int64        //读操作需要用到的偏移量
	wmutex  sync.Mutex   //写操作需要用到的互斥锁
	rmutex  sync.Mutex   //读操作需要用到的互斥锁
	dataLen uint32       //数据块长度
}

// 初始化DataFile类型值的函数，返回一个DataFile类型的值
func NewDataFile(path string, dataLen uint32) (DataFile, error) {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Fail to find", f, "cServer start Failed")
		return nil, err
	}

	if dataLen == 0 {
		return nil, errors.New("Invalid data length!")
	}

	df := &myDataFile{
		f:       f,
		dataLen: dataLen,
	}

	return df, nil
}

// 获取并更新读偏移量，根据读偏移量从文件中读取一块数据，把该数据块封装成一个Data类型值并将其结果返回
func (df *myDataFile) Read(done <-chan bool) (rsn int64, d Data, err error) {
	// 读取并更新读偏移量
	var offset int64
	// 读互斥锁定
	df.rmutex.Lock()
	offset = df.roffset
	// 更改偏移量，当前偏移量+数据块长度
	df.roffset += int64(df.dataLen)
	// 读互斥解锁
	df.rmutex.Unlock()

	// 读取一个数据块，最后读取的数据块序列号
	rsn = offset / int64(df.dataLen)
	bytes := make([]byte, df.dataLen)
	for {
		// 读写锁：读锁定
		df.fmutex.RLock()
		_, err := df.f.ReadAt(bytes, offset)
		// fmt.Println("offset: ", offset, ", bytes:", bytes)

		if err != nil {
			// 由于进行写操作的Groutine比进行读操作的Goroutine少，所以过不了多久读偏移量roffset的值会大于写偏移量woffset的值
			// 也就是说，读操作很快就没有数据块可读了，这种情况会让df.f.ReadAt方法返回的第二个结果值为代表的非nil且会与io.EOF相等的值
			// 因此，不应该把EOF看成错误的边界情况
			// So 在读操作读完数据块，EOF时解锁读操作，并继续循环，尝试获取同一个数据块，直到获取成功为止。
			if err == io.EOF {
				// 如果写入结束，则退出Goroutine
				if len(done) == 3 {
					break
				}
				// 注意，如果在该for代码块被执行期间，一直让读写锁fmutex处于读锁定状态，那么针对它的写操作将永远不会成功。
				// 且相应的Goroutine也会被一直堵塞，因为它们是互斥的。
				// so 在每条return & continue语句的前面加入一个针对读写锁的解锁操作
				df.fmutex.RUnlock()
				// 注意，出现EOF时可能是很多意外情况，比如文件被删除，文件损坏等
				// 这里可以考虑把逻辑提交给上层处理。
				continue
			}
		}
		break
	}
	d = bytes
	df.fmutex.RUnlock()
	return
}

func (df *myDataFile) Write(d Data) (wsn int64, err error) {
	// 读取并更新写偏移量
	var offset int64
	df.wmutex.Lock()
	offset = df.woffset
	df.woffset += int64(df.dataLen)
	df.wmutex.Unlock()

	// 写入一个数据块，最后写入数据块的序号
	wsn = offset / int64(df.dataLen)
	var bytes []byte
	if len(d) > int(df.dataLen) {
		bytes = d[0:df.dataLen]
	} else {
		bytes = d
	}
	df.fmutex.Lock()
	_, err = df.f.Write(bytes)
	df.fmutex.Unlock()

	return
}

func (df *myDataFile) Rsn() int64 {
	df.rmutex.Lock()
	defer df.rmutex.Unlock()
	return df.roffset / int64(df.dataLen)
}

func (df *myDataFile) Wsn() int64 {
	df.wmutex.Lock()
	defer df.wmutex.Unlock()
	return df.woffset / int64(df.dataLen)
}

func (df *myDataFile) DataLen() uint32 {
	return df.dataLen
}

func main() {
	var dataFile DataFile
	dataFile, _ = NewDataFile("./mutex_2015_1.dat", 20)

	var d = map[int]Data{
		1: []byte("bata_test1\n"),
		2: []byte("batu_test2\n"),
		3: []byte("test1_batu\n"),
	}

	done := make(chan bool, 3)

	//写入数据
	for i := 1; i < 4; i++ {
		go func(i int) {
			wsn, _ := dataFile.Write(d[i])
			fmt.Println("write i=", i, ", wsn=", wsn, ", success.")
			done <- true
		}(i)
	}

	//读取数据
	for i := 1; i < 4; i++ {
		go func(i int) {
			rsn, d, _ := dataFile.Read(done)
			fmt.Printf("Read i=%d, rsn=%d, data=%q, success.\n", i, rsn, string(d))
		}(i)
	}

	time.Sleep(10 * time.Second)
}
