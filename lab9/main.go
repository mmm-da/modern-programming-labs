package main

// Телнет. Создать программу, которая соединяется с указанным
// сервером по указанному порту и производит обмен текстовой информацией.

import (
	"net"
	"os"
)

func readStdin(stop chan struct{}) chan []byte {
	bytesFromStdin := make(chan []byte)
	go func() {
		defer close(bytesFromStdin)
		buffer := make([]byte, 10)
		for {
			select {
			case <-stop:
				break
			default:
				os.Stdin.Read(buffer)
				bytesFromStdin <- buffer
			}
		}
	}()
	return bytesFromStdin
}

func readTelnet(telnetSocket net.Conn, stop chan struct{}) chan []byte {
	bytesFromTelnet := make(chan []byte)

	go func() {
		defer close(bytesFromTelnet)
		buffer := make([]byte, 10)
		for {
			select {
			case <-stop:
				break
			default:
				telnetSocket.Read(buffer)
				bytesFromTelnet <- buffer
			}
		}
	}()
	return bytesFromTelnet
}

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:2701")
	if err != nil {
		panic(err)
	}

	stop := make(chan struct{})
	defer close(stop)
	stdin := readStdin(stop)
	telnet := readTelnet(conn, stop)

	for {
		select {
		case char := <-stdin:
			conn.Write(char)
		case char := <-telnet:
			os.Stdout.Write(char)
		}
	}
}
