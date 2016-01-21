package main

//net dial
import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()
	copy(os.Stdout, conn)
}

func copy(w io.Writer, r io.Reader) {
	if _, err := io.Copy(w, r); err != nil {
		log.Fatal(err)
	}
}
