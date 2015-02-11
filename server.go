package main
import "net"
import "log"

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn,err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
	}
	defer ln.Close()
}
