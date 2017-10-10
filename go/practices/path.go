package main

import (
	"fmt"
	"path"
)

var templates = []string{
	path.Join("/var/www/html", "templates", "index.html"),
	path.Join("templates", "widgets", "login.html"),
	path.Join("templates", "widgets", "header.html"),
	path.Join("templates", "widgets", "footer.html"),
	path.Join("templates", "widgets", "topbar.html"),
	path.Join("templates", "widgets", "left-menu.html"),
	path.Join("templates", "projects.html"),
}

func main() {
	p := path.Join("templates", "index.html")
	fmt.Println(p)

	fmt.Println(templates)
}
