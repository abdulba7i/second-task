package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

func main() {
	var timeout int
	flag.IntVar(&timeout, "timeout", 10, "timeout in seconds")
	flag.Parse()

	args := flag.Args()
	if len(args) != 2 {
		fmt.Println("Usage: go run main.go [--timeout=10] host port")
		os.Exit(1)
	}

	host := args[0]
	port := args[1]

	address := fmt.Sprintf("%s:%s", host, port)

	conn, err := net.DialTimeout("tcp", address, time.Duration(timeout)*time.Second)
	if err != nil {
		fmt.Printf("Error connecting to %s: %v\n", address, err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Printf("Connected to %s\n", address)

	done := make(chan bool)

	go func() {
		reader := bufio.NewReader(conn)
		for {
			data, err := reader.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					fmt.Println("Server closed connection")
				} else {
					fmt.Printf("Error reading from server: %v\n", err)
				}
				done <- true
				return
			}
			fmt.Print(data)
		}
	}()

	go func() {
		reader := bufio.NewReader(os.Stdin)
		for {
			data, err := reader.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					fmt.Println("Ctrl+D pressed, closing connection")
				} else {
					fmt.Printf("Error reading from stdin: %v\n", err)
				}
				done <- true
				return
			}
			_, err = conn.Write([]byte(data))
			if err != nil {
				fmt.Printf("Error writing to server: %v\n", err)
				done <- true
				return
			}
		}
	}()

	<-done
	fmt.Println("Connection closed")
}
