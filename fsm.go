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
    case strings.HasPrefix(line, "/about"):
        return fsm_about, line[6:], nil
	}
	return fsm_heuristic, line, nil
}

func fsm_help(conn io.ReadWriter, log io.Writer, args string) (next transition, rest string, err error) {
	conn.Write([]byte(line_to_print("Available commands:")))
	conn.Write([]byte(line_to_print("/help              show this message")))
	conn.Write([]byte(line_to_print("/msg [content]     leave " + USER_NAME + " a message")))
	conn.Write([]byte(line_to_print("/about             something about me, the robot")))
	return fsm_start, "", nil
}

func fsm_about(conn io.ReadWriter, log io.Writer, args string) (next transition, rest string, err error) {
    conn.Write([]byte(line_to_print("Hey! I'm open sourced on https://github.com/songgao/hey")))
    return fsm_start, "", nil
}

func fsm_heuristic(conn io.ReadWriter, log io.Writer, args string) (next transition, rest string, err error) {
	conn.Write([]byte(line_to_print("Heuristic not implemented. Falling back to /help.")))
	return fsm_help, args, nil
}

func fsm_msg(conn io.ReadWriter, log io.Writer, args string) (next transition, rest string, err error) {
	_, err = fmt.Fprintf(log, "Message: %s\n", args)
	if err == nil {
		conn.Write([]byte(line_to_print("Sure. I'll let " + correct_gender("him") + " know.")))
	} else {
		conn.Write([]byte(line_to_print("Something's wrong; " + correct_gender("he") + "'s probably not gonna get the message. Sorry!")))
	}
	return fsm_start, "", err
}
