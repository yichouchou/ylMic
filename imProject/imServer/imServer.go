package main

import (
	"log"
	"net"
	"time"
)

const (
	SERVER_IP   = "127.0.0.1"
	SERVER_PORT = "120"
)

type Server struct {
}

func (server *Server) Start() {

	serverAddr := SERVER_IP + ":" + SERVER_PORT
	listener, err := net.Listen("tcp", serverAddr)
	if err != nil {
		log.Println("Failed to execute function net.Listen: ", err.Error())
		return
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			//下方是处理连接异常的日志报错
			if nerr, ok := err.(net.Error); ok && nerr.Temporary() {
				log.Println("Continue accepting connection because of an temporary error: ", err.Error())
				time.Sleep(1 * time.Second)
				continue
			}

			log.Println("Failed to execute function listener.Accept, exit app because of an error: ", err.Error())
			return
		}

		sess := &Session{
			Conn: conn,
		}
		go sess.Serve()
	}
}

func main() {
	(&Server{}).Start()
}
