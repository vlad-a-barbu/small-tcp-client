package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Println("Invalid args. Please provide the server address and a port number.")
		return
	}

	serverAddress := os.Args[1]
	port := os.Args[2]

	connectionString := fmt.Sprintf("%s:%s", serverAddress, port)

	conn, err := net.Dial("tcp", connectionString)

	if err != nil {
		fmt.Println(err)
		return
	}

	closingConnSuffix := "closing connection\n"

	for {
		clientReader := bufio.NewReader(os.Stdin)
		serverReader := bufio.NewReader(conn)

		serverResponse, err := serverReader.ReadString('\n')

		if strings.HasSuffix(serverResponse, closingConnSuffix) {
			log.Fatalln("The server closed your connection")
		}

		switch err {
		case io.EOF:
			{
				fmt.Println("\nThe server is either down or closed your connection")
				return
			}
		case nil:
			{
				fmt.Println("\nServer response:")
				fmt.Printf("> %s\n", serverResponse)
				for {
					fmt.Printf("Request:\n> ")
					request, _ := clientReader.ReadString('\n')
					if len(request) > 1 {
						conn.Write([]byte(request))
						break
					}
				}
			}
		default:
			{
				fmt.Println(err)
				return
			}
		}
	}
}
