package main

// Телнет. Создать программу, которая соединяется с указанным
// сервером по указанному порту и производит обмен текстовой информацией.

import (
	"io"
	"net"
	"os"
)

func connectStreams(dst io.Writer, src io.Reader, stop chan struct{}) {
	go func() {
		for {
			select {
			case <-stop:
				break
			default:
				io.Copy(dst, src)
			}
		}
	}()
}

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:2701")
	if err != nil {
		panic(err)
	}

	stop := make(chan struct{})
	defer close(stop)
	connectStreams(conn, os.Stdin, stop)
	connectStreams(os.Stdout, conn, stop)
	for {
	}
}
