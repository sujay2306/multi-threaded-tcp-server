package main

import (
	"log"
	"net"
	"time"
)


func do(conn net.Conn){

	buf := make ([] byte, 1024)
 	_,err := conn.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("processing the request")
	time.Sleep(8 * time.Second) //wait for 1 sec
	conn.Write([] byte ("HTTP/1.1 200 OK\r\n\r\nHello I'm TCP Server!!!\r\n"))
	conn.Close()
}

func main(){
	listener, err := net.Listen("tcp", ":1729")
	if err != nil {
		log.Fatal(err)
	}
	
	for{
		log.Println("waiting for client connection")
		conn,err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		log.Println("client connected")
		go do(conn)

	}

}
