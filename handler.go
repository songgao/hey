package main

import (
	"net"
	"os"
	"path"
	"time"
)

func handle(conn net.Conn) {
	conn.Write(TEXT_greeting)
	file, err := os.Create(path.Join(logDir, time.Now().Format(time.RFC3339)+"-"+conn.RemoteAddr().String()))
	var next transition = fsm_start
	var args string = ""
	for err == nil {
		next, args, err = next(conn, file, args)
	}
	if conn != nil {
		conn.Close()
	}
	if file != nil {
		file.Close()
	}
}
