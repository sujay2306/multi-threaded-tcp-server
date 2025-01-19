package main

import (
	"fmt"
	"log"
	"net"
)


func do(conn, net.Conn){
	buf := make ([] byte, 1024)
 	conn.Read(buf)
	conn.Write()
	conn.Close()
}
func main(){
	listener, err := net.Listen("tcp", ":1729")
	if err != nil {
		log.Fatal(err)
	}
	
	conn , err := listener.Accept()
	if err != nil {
		log.Fatal(err)
	}


}