package main

import (
	"fmt"
	"strings"
)

func main() {
	s := strings.TrimLeft("/srv/salt/scripts/immortal-upgrade.sh", "/srv/salt/")
	a := "/srv/salt/scripts/immortal-upgrade.sh"[len("/srv/salt/"):]
	fmt.Println(s, a)
}
