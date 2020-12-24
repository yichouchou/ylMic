package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

const (
	//根据自身情况修改服务器ip地址
	serverAddr = "127.0.0.1:120"
)

func main() {
	succCount := 0
	failCount := 0
	for i := 0; i < 60000; i++ {
		conn, err := net.Dial("tcp", serverAddr)
		if err != nil {
			failCount++
			fmt.Println("Number of dialing failures:", failCount)
			time.Sleep(time.Second)
			continue
		}
		defer conn.Close()

		fmt.Println("Number of successful connections:", succCount)

		go Client(conn, succCount)

		succCount++
	}

	//Stop here to prevent the program from exiting
	tick := time.NewTicker(10 * time.Second)
	for {
		select {
		case <-tick.C:
		}
	}
}

func Client(conn net.Conn, index int) error {
	dataRead := make([]byte, 512)
	dataWrite := strconv.Itoa(index) +
		"1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890\n" +
		"1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890\n" +
		"1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890\n" +
		"1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890\n" +
		"1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890\n"

	for {
		conn.Write([]byte(dataWrite))
		_, err := conn.Read(dataRead)
		if err != nil {
			return err
		}
		time.Sleep(60 * time.Second)
	}
	return nil
}
