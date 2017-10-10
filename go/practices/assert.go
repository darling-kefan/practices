package main

import "fmt"

type Mover interface {
	Move()
}

type Locker interface {
	Lock()
	Unlock()
}

type MoveLocker interface {
	Mover
	Locker
}

type bike struct{}

func (bike) Move() {
	fmt.Println("Moving the bike")
}

func (bike) Lock() {
	fmt.Println("Locking the bike")
}

func (bike) Unlock() {
	fmt.Println("Unlocking the bike")
}

func main() {
	var m Mover
	var ml MoveLocker

	ml = bike{}
	m = ml

	m.Move()
	ml.Move()
	ml.Lock()
	ml.Unlock()

	b := m.(bike)
	ml = b

	ml.Move()
	ml.Lock()
	ml.Unlock()
}
