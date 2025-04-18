package main

import (
	"fmt"
	"syscall"
)

func main() {
	//? Create a socket
	serverFd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	if err != nil {
		fmt.Println("Socket error:", err)
		return
	}
	defer syscall.Close(serverFd)
	syscall.SetsockoptInt(serverFd, syscall.SOL_SOCKET, syscall.SO_REUSEADDR, 1)

	// Bind socket to address and port
	addr := syscall.SockaddrInet4{Port: 8080}
	copy(addr.Addr[:], []byte{127, 0, 0, 1}) // localhost

	err = syscall.Bind(serverFd, &addr)
	if err != nil {
		fmt.Println("Bind error:", err)
		return
	}

	// Start listening
	err = syscall.Listen(serverFd, 10)
	if err != nil {
		fmt.Println("Listen error:", err)
		return
	}

	fmt.Println("Server is running on http://127.0.0.1:8080")

	for {
		// Accept a new connection
		clientFd, _, err := syscall.Accept(serverFd)
		if err != nil {
			fmt.Println("Accept error:", err)
			continue
		}

		// Read request
		buf := make([]byte, 1024)
		_, err = syscall.Read(clientFd, buf)
		if err != nil {
			fmt.Println("Read error:", err)
			syscall.Close(clientFd)
			continue
		}

		// HTTP Response
		response := "HTTP/1.1 200 OK\r\n" +
			"Content-Type: text/plain\r\n" +
			"Content-Length: 13\r\n" +
			"\r\n" +
			"Hello, World!"

		// Send response
		_, err = syscall.Write(clientFd, []byte(response))
		if err != nil {
			fmt.Println("Write error:", err)
		}

		// Close client connection
		syscall.Close(clientFd)
	}
}
