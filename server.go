package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
)

type Message struct {
	sender  int
	message string
}

func handleError(err error) {
	// TODO: all
	// Deal with an error event.
	return
}

func acceptConns(ln net.Listener, conns chan net.Conn) {
	// TODO: all
	// Continuously accept a network connection from the Listener
	// and add it to the channel for handling connections.
	for {
		conn, _ := ln.Accept()
		conns <- conn
	}

}

func handleClient(client net.Conn, clientid int, msgs chan Message) {
	reader := bufio.NewReader(client)
	// TODO: all
	// So long as this connection is alive:
	// Read in new messages as delimited by '\n's
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			//fmt.Println("I AM HERE!")
			handleError(err)
		} else {
			msgs <- Message{clientid, msg}
		}
		// Tidy up each message and add it to the messages channel,
		// recording which client it came from.
	}
}

func main() {
	// Read in the network port we should listen on, from the commandline argument.
	// Default to port 8030
	portPtr := flag.String("port", ":8030", "port to listen on")
	flag.Parse()

	//TODO Create a Listener for TCP connections on the port given above.
	ln, _ := net.Listen("tcp", *portPtr)
	//Create a channel for connections
	conns := make(chan net.Conn)
	//Create a channel for messages
	msgs := make(chan Message)
	//Create a mapping of IDs to connections
	clients := make(map[int]net.Conn) //empty mao - need to assign key and value ourselves

	//Start accepting connections
	go acceptConns(ln, conns)
	for {
		select {
		case conn := <-conns:
			//TODO Deal with a new connection
			// - assign a client ID
			idClient := len(clients) + 1
			clients[idClient] = conn //add new key to our map (assigning key to conn as well)
			// - add the client to the clients channel
			go handleClient(conn, idClient, msgs)
			// does it get added to the end of queue and become selected ??
			// - start to asynchronously handle messages from this client

		case msg := <-msgs:
			//TODO Deal with a new message
			// Send the message to all clients that aren't the sender

			for clientid, client := range clients {
				if clientid != msg.sender {
					fmt.Fprintln(client, msg.message)
				}
			}
		default:

		}
	}
}
