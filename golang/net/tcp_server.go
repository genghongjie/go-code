package main

import (
	"log"
	"net"
)

func main() {
	// Part 1: create a listener
	l, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("Error listener returned: %s", err)
	}
	defer l.Close()

	for {
		// Part 2: accept new connection
		c, err := l.Accept()
		if err != nil {
			log.Fatalf("Error to accept new connection: %s", err)
		}

		// Part 3: create a goroutine that reads and write back data
		go func() {
			log.Printf("TCP session open")
			defer c.Close()

			for {
				d := make([]byte, 100)

				// Read from TCP buffer
				_, err := c.Read(d)
				if err != nil {
					log.Printf("Error reading TCP session: %s", err)
					break
				}
				log.Printf("reading data from client: %s\n", string(d))

				// write back data to TCP client
				_, err = c.Write(d)
				if err != nil {
					log.Printf("Error writing TCP session: %s", err)
					break
				}
			}
		}()
	}
}
