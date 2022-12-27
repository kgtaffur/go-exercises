package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {
	p := fmt.Println
	s := "Mysql://admin:password@serverhost.com:8081/server/page1?key=value&key2=value2#X"

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	p("URL:", s)
	p("Scheme:", u.Scheme)
	p("User:", u.User)
	p("Username:", u.User.Username())

	pass, _ := u.User.Password()
	p("Password:", pass)
	p("Host and port:", u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	p("Host:", host)
	p("Port", port)
	p("Path:", u.Path)
	p("Fragment:", u.Fragment)
	p("RawQuery:", u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	p("Parsed RawQuery:", m)
	p("Value from position 0 with key 'key2' in the map:", m["key2"][0])
}
