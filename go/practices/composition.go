package main

import (
	"fmt"
)

type administrator interface {
	administrate(system string)
}

type developer interface {
	develop(system string)
}

// ========================================================================

type adminlist struct {
	list []administrator
}

func (l *adminlist) Enqueue(a administrator) {
	l.list = append(l.list, a)
}

func (l *adminlist) Dequeue() administrator {
	a := l.list[0]
	l.list = l.list[1:]
	return a
}

// ========================================================================

type devlist struct {
	list []developer
}

func (l *devlist) Enqueue(d developer) {
	l.list = append(l.list, d)
}

func (l *devlist) Dequeue() developer {
	d := l.list[0]
	l.list = l.list[1:]
	return d
}

// =======================================================================

type sysadmin struct {
	name string
}

func (s *sysadmin) administrate(system string) {
	fmt.Println("administrator: ", s.name, system)
}

type programmer struct {
	name string
}

func (p *programmer) develop(system string) {
	fmt.Println("developer: ", p.name, system)
}

type compony struct {
	administrator
	developer
}

// ========================================================================

func main() {
	var admins adminlist
	var devs devlist

	admins.Enqueue(&sysadmin{name: "shouqiang"})
	devs.Enqueue(&programmer{name: "kefan"})
	devs.Enqueue(&programmer{name: "chunhuan"})

	cmp := compony{
		administrator: admins.Dequeue(),
		developer:     devs.Dequeue(),
	}

	admins.Enqueue(cmp)
	devs.Enqueue(cmp)

	tasks := []struct {
		needsAdmin bool
		system     string
	}{
		{needsAdmin: false, system: "xenia"},
		{needsAdmin: true, system: "pillar"},
		{needsAdmin: false, system: "omega"},
	}

	fmt.Println(len(admins.list), len(devs.list))

	// Iterate over tasks
	for _, task := range tasks {
		if task.needsAdmin {
			admins.Dequeue().administrate(task.system)
			continue
		}

		devs.Dequeue().develop(task.system)
	}

	fmt.Println(len(admins.list), len(devs.list))
}
