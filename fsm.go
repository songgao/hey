package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type transition func(conn io.ReadWriter, log io.Writer, args string) (next transition, rest string, err error)

func fsm_start(conn io.ReadWriter, log io.Writer, args string) (next transition, rest string, err error) {
	line, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		return fsm_start, "", err
	}
	line = strings.Trim(line, " ")
	switch {
	case strings.HasPrefix(line, "/help"):
		return fsm_help, line[5:], nil
	case strings.HasPrefix(line, "/msg "):
		return fsm_msg, line[5:], nil
	}
	return fsm_heuristic, line, nil
}

func fsm_help(conn io.ReadWriter, log io.Writer, args string) (next transition, rest string, err error) {
	conn.Write([]byte(line_to_print("Available commands:")))
	conn.Write([]byte(line_to_print("/help              show this message")))
	conn.Write([]byte(line_to_print("/msg [content]     leave " + USER_NAME + " a message")))
	return fsm_start, "", nil
}

func fsm_heuristic(conn io.ReadWriter, log io.Writer, args string) (next transition, rest string, err error) {
	conn.Write([]byte(line_to_print("Heuristic not implemented. Falling back to /help.")))
	return fsm_help, args, nil
}

func fsm_msg(conn io.ReadWriter, log io.Writer, args string) (next transition, rest string, err error) {
	fmt.Fprintf(log, "Message: %s\n", args)
	conn.Write([]byte(line_to_print("Sure. I'll let " + correct_gender("him") + " know.")))
	return fsm_start, "", nil
}
