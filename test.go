package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {

	ln, err := net.Listen("tcp", "localhost:1234")
	if err != nil {
		log.Fatalln(err)
	}

	go monitor()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConnection(conn)
	}
}

var (
	message  = make(chan string)
	entering = make(chan client)
	leaving  = make(chan client)
)

type client chan<- string

func monitor() {
	clients := make(map[client]bool)

	for {
		select {
		case ch := <-entering:
			clients[ch] = true
		case ch := <-leaving:
			delete(clients, ch)
		case msg := <-message:
			for ch := range clients {
				ch <- msg // opti
			}
		}
	}
}

func handleConnection(conn net.Conn) {
	ch := make(chan string)
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	msg := fmt.Sprintf("u r :%s , welcome!\r\n", who)
	ch <- msg

	message <- who + " has arrived!"
	entering <- ch

	input := bufio.NewScanner(conn)
	for input.Scan() {
		message <- who + ": " + input.Text()
	}

	leaving <- ch
	message <- who + " has left!"

	close(ch)
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}
