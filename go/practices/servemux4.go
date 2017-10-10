// DefaultServerMux就是我们之前用到的ServerMux, 只是它随着net/http包初始化的
// 时候被自动初始化了而已。
//
//   var DefaultServeMux = NewServeMux()
//
// net/http包提供了一组快捷方式来配合DefaultServeMux: http.Handle和http.HandleFunc。
// 这些函数与我们之前看过的类似的名称的函数功能一样，唯一的不同是他们将处理器注册到DefaultServerMux，而之前我们是注册到自己创建的ServeMux
//
// 此外，ListenAndServe在没有提供其他的处理器的情况下(也就是第二个参数设成了nil)，内部会使用DefaultServeMux。

package main

import (
	"log"
	"net/http"
	"time"
)

func timeHandler(format string) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		tm := time.Now().Format(format)
		w.Write([]byte("The time is: " + tm))
	}
	return http.HandlerFunc(fn)
}

func main() {
	var format string = time.RFC1123
	th := timeHandler(format)

	// We use http.Handle instead of mux.Handle...
	http.Handle("/time", th)

	log.Println("Listening...")
	// And pass nil as the handler to ListenAndServe
	http.ListenAndServe(":3000", nil)
}
