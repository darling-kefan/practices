package main

type nexter interface {
	next() byte
}

func nextFew1(n nexter, num int) []byte {
	var b []byte
	for i := 0; i < num; i++ {
		b[i] = n.next()
	}
	return b
}

func nextFew2(n *nexter, num int) []byte {
	var b []byte
	for i := 0; i < num; i++ {
		b[i] = n.next()
	}
	return b
}

func main() {

}
