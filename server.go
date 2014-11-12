package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		accept(conn)
		defer conn.Close()
	}
}

func accept(c net.Conn) {
	b := make([]byte, 1024)
	for {
		_, err := c.Read(b)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if err == io.EOF {
			return
		}
		fmt.Println(string(b))
		fmt.Println("====")
		c.Write([]byte("aaaaa\nbbbbb\nccccccc\nddddddd\n\n\n\n\n\n\n\n"))
		c.Close()
		return
	}
}
