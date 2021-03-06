package main

import (
	"io"
	"log"
	"net"
	"os"
)

func copyMessage(dst io.Writer, src io.Reader) {
	_, err := io.Copy(dst, src)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()
	copyMessage(os.Stdout, conn)

}
