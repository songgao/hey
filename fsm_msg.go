package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

func fsm_msg(conn io.ReadWriter, log io.Writer, args string) (next transition, rest string, err error) {
	if !strings.HasSuffix(args, "\n") {
		args = args + "\n"
	}
	_, err = fmt.Fprintf(log, "Message: %s", args)
	if err == nil {
		conn.Write(TEXT_message_logged)
	} else {
		conn.Write(TEXT_message_failed)
	}
	return fsm_start, "", err
}

func fsm_msgstart(conn io.ReadWriter, log io.Writer, args string) (next transition, rest string, err error) {
	_, err = fmt.Fprintf(log, "<======== Multi-line Message Start ========>\n%s", args)
	if err == nil {
		conn.Write(TEXT_message_start)
	} else {
		conn.Write(TEXT_message_failed)
	}
	reader := bufio.NewReader(conn)
	if !strings.Contains(args, "/msgend") {
		for err == nil {
			line, err := reader.ReadSlice('\n')
			_, err2 := log.Write(line)
			if err != nil || err2 != nil {
				conn.Write(TEXT_message_failed)
				return fsm_start, "", err
			}
			if bytes.Contains(line, []byte("/msgend")) {
				break
			}
		}
	}
	conn.Write(TEXT_message_logged)
	fmt.Fprintf(log, "<========  Multi-line Message End  ========>\n")
	return fsm_start, "", err
}
