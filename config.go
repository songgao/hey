package main

import (
	"os"
)

const (
	TLS_LISTEN  = ":56789"
	TCP_LISTEN  = ":56788"
	ROBOT_NAME  = "Oag Gnos"
	USER_NAME   = "Song Gao"
	USER_GENDER = MALE
)

var (
	keyFile  = os.Getenv("HOME") + "/.hey/server.key"
	certFile = os.Getenv("HOME") + "/.hey/server.pem"
	logDir   = os.Getenv("HOME") + "/.hey/logs/"
)

const (
	_ = iota
	MALE
	FEMALE
)
