package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type Session struct {
	Conn   net.Conn
	Reader *bufio.Reader
	Writer *bufio.Writer
}

func (session *Session) Serve() {
	defer session.Close()

	session.Reader = bufio.NewReader(session.Conn)
	session.Writer = bufio.NewWriter(session.Conn)

	for {
		byteData, err := session.Reader.ReadSlice('\n')
		if err != nil {
			log.Println("Error occured in funciton Session.Serve: ", err.Error())
			return
		}

		line := string(byteData)
		fmt.Println(line)
		session.Send(line)
	}
}

func (session *Session) Close() {
	session.Conn.Close()
}

func (session *Session) write(data string) error {
	if _, err := fmt.Fprint(session.Writer, data); err != nil {
		return err
	}
	return session.Writer.Flush()
}

func (session *Session) Send(data string) error {
	return session.write(data)
}
