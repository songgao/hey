package main

import (
	"crypto/tls"
	"fmt"
	"net"
	"os"
	"os/exec"
	"sync"
)

func init() {
	os.MkdirAll(logDir, 0700)
}

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(2)
	go listen_tls(wg)
	go listen_tcp(wg)
	wg.Wait()
}

func listen_tcp(wg *sync.WaitGroup) {
	ln, err := net.Listen("tcp", TCP_LISTEN)
	if err != nil {
		fmt.Printf("Error in listening TCP: %v\n", err)
	}
	for {
		conn, err := ln.Accept()
		if err == nil {
			go handle(conn)
		} else {
			fmt.Println(err)
		}
	}
	wg.Done()
}

func listen_tls(wg *sync.WaitGroup) {
	cert, ok := getCert()
	if !ok {
		fmt.Printf("\nFailed to load x509 certificate.\n")
		return
	}
	config := new(tls.Config)
	config.Certificates = []tls.Certificate{cert}
	ln, err := tls.Listen("tcp", TLS_LISTEN, config)
	if err != nil {
		fmt.Printf("Error creating listener: %v\n", err)
		return
	}
	for {
		conn, err := ln.Accept()
		if err == nil {
			go handle(conn)
		} else {
			fmt.Println(err)
		}
	}
	wg.Done()
}

func getCert() (tls.Certificate, bool) {
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		fmt.Println("Certificate is not found or cannot be loaded.")
		fmt.Println(certFile)
		fmt.Println(keyFile)
		fmt.Print("Would you like to generate a \"whatever\" one? (y/n): ")
		var yn string
		fmt.Scanln(&yn)
		if yn != "y" && yn != "Y" {
			return cert, false
		}
		err = exec.Command("openssl", "req", "-new", "-nodes", "-x509", "-out", certFile, "-keyout", keyFile, "-days", "4096", "-subj", `/C=/ST=/L=/O=/OU=/CN=/emailAddress=email@example.com`).Run()
		if err != nil {
			return cert, false
		}
		cert, _ = tls.LoadX509KeyPair(certFile, keyFile)
	}
	return cert, true
}
