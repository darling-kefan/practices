package main

import (
	"fmt"
	_ "os"
	_ "runtime"
	"sync"
)

//func main() {
//    runtime.GOMAXPROCS(runtime.NumCPU())
//
//    chan_n := make(chan bool)
//    chan_c := make(chan bool) // 等价于chan_c := make(chan bool, 0)
//
//    var wg sync.WaitGroup
//    wg.Add(2)
//
//    go func() {
//        for i := 1; i <= 9; i += 2 {
//            <-chan_c
//            fmt.Print(i)
//            fmt.Print(i + 1)
//            if i == 9 {
//                close(chan_n)
//            } else {
//                chan_n <- true
//            }
//        }
//        wg.Done()
//    }()
//
//    go func() {
//        char_seq := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K"}
//        for i := 0; i <= 8; i += 2 {
//            _, ok := <-chan_n
//            fmt.Print(char_seq[i])
//            fmt.Print(char_seq[i+1])
//            if ok {
//                chan_c <- true
//            }
//        }
//        wg.Done()
//    }()
//
//    chan_c <- true
//    wg.Wait()
//}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	chan_n := make(chan bool, 1)
	chan_c := make(chan bool)

	go func() {
		for i := 1; i <= 9; i += 2 {
			<-chan_n
			fmt.Print(i)
			fmt.Print(i + 1)
			chan_c <- true
		}
		wg.Done()
	}()

	go func() {
		letters := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}
		for i := 0; i <= 8; i += 2 {
			<-chan_c
			fmt.Print(letters[i])
			fmt.Print(letters[i+1])
			chan_n <- true
		}
		wg.Done()
	}()

	chan_n <- true
	wg.Wait()
}
