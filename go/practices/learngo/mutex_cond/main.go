package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

type DataFile interface {
	Read() (rsn int64, d Data, err error)
	Write(d Data) (wsn int64, err error)
	Rsn() int64
	Wsn() int64
	DataLen() uint32
}

type Data []byte

type myDataFile struct {
	f       *os.File
	fmutex  sync.RWMutex
	rcond   *sync.Cond //读操作时需要用到的条件变量
	woffset int64
	roffset int64
	wmutex  sync.Mutex
	rmutex  sync.Mutex
	dataLen uint32
}

func NewDataFile(path string, dataLen uint32) (DataFile, error) {
	f, err := os.Create(path)
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

	//创建一个可用的条件变量(初始化),返回一个*sync.Cond类型的结果值，我们就可以调用该值拥有的三个方法Wait,Signal,Broadcast
	df.rcond = sync.NewCond(df.fmutex.RLocker())
	return df, nil
}

func (df *myDataFile) Read() (rsn int64, d Data, err error) {
	var offset int64
	df.rmutex.Lock()
	offset = df.roffset
	df.roffset += int64(df.dataLen)
	df.rmutex.Unlock()

	rsn = offset / int64(df.dataLen)
	bytes := make([]byte, df.dataLen)
	df.fmutex.RLock()
	defer df.fmutex.RUnlock()

	for {
		_, err = df.f.ReadAt(bytes, offset)
		if err != nil {
			if err == io.EOF {
				// 暂时放弃fmutex的读锁，并等待通知的到来
				df.rcond.Wait()
				continue
			}
		}
		break
	}
	d = bytes
	return
}

func (df *myDataFile) Write(d Data) (wsn int64, err error) {
	var offset int64
	df.wmutex.Lock()
	offset = df.woffset
	df.woffset += int64(df.dataLen)
	df.wmutex.Unlock()

	wsn = offset / int64(df.dataLen)
	var bytes []byte
	if len(d) > int(df.dataLen) {
		bytes = d[0:df.dataLen]
	} else {
		bytes = d
	}
	df.fmutex.Lock()
	defer df.fmutex.Unlock()
	_, err = df.f.Write(bytes)
	//发送通知
	df.rcond.Signal()
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
	dataFile, _ = NewDataFile("./mutex_2915_1.dat", 10)

	var d = map[int]Data{
		1: []byte("batu_test1"),
		2: []byte("batu_test2"),
		3: []byte("batu_test3"),
		4: []byte("batu_3"),
	}

	for i := 1; i < 4; i++ {
		go func(i int) {
			wsn, _ := dataFile.Write(d[i])
			fmt.Println("write i=", i, ",wsn=", wsn, ",success.")
		}(i)
	}

	for i := 1; i < 5; i++ {
		go func(i int) {
			rsn, d, _ := dataFile.Read()
			fmt.Println("Read i=", i, ",rsn=", rsn, ",data=", string(d), ",success.")
		}(i)
	}

	time.Sleep(10 * time.Second)
}
