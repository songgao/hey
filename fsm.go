package main

import (
	"bufio"
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
	case strings.HasPrefix(line, "/about"):
		return fsm_about, line[6:], nil
	case strings.HasPrefix(line, "/msg "):
		return fsm_msg, line[5:], nil
	case strings.HasPrefix(line, "/msgstart"):
		return fsm_msgstart, line, nil
	}
	return fsm_heuristic, line, nil
}

func fsm_help(conn io.ReadWriter, log io.Writer, args string) (next transition, rest string, err error) {
	conn.Write(TEXT_help)
	return fsm_start, "", nil
}

func fsm_about(conn io.ReadWriter, log io.Writer, args string) (next transition, rest string, err error) {
	conn.Write(TEXT_about)
	return fsm_start, "", nil
}

func fsm_heuristic(conn io.ReadWriter, log io.Writer, args string) (next transition, rest string, err error) {
	conn.Write(TEXT_heuristic_not_implemented)
	return fsm_help, args, nil
}
