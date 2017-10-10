package main

type Server interface {
	Start() error
	Wait() error
	Stop() error
}

type server struct {
	host string
}

func NewServer(host string) Server {
	return &server{host: host}
}

/*
type Server struct {
	host string
}

func NewServer(host string) *Server {
	return &Server{host: host}
}
*/

func (s *server) Start() error {
	return nil
}

func (s *server) Stop() error {
	return nil
}

func (s *server) Wait() error {
	return nil
}

func main() {
	s := NewServer("127.0.0.1")

	s.Start()
	s.Wait()
	s.Stop()
}
