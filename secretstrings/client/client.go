package main

import (
	"bufio"
	"strings"
	"uk.ac.bris.cs/distributed2/secretstrings/stubs"

	//	"net/rpc"
	"flag"
	"log"
	"net/rpc"
	"os"
	//	"bufio"
	//	"os"
	//	"uk.ac.bris.cs/distributed2/secretstrings/stubs"
	"fmt"
)

//func aquireLock(c chan int) {
//	_ = <-c
//}
//
//func releaseLock(c chan int) {
//	c <- 0
//}

func getWords(clients []*rpc.Client) {
	response := new(stubs.Response)
	//var serverLocks []chan int
	//for range clients {
	//	lock := make(chan int, 1)
	//	serverLocks = append(serverLocks, lock)
	//}

	file, err := os.Open("/Users/davidpayne/Desktop/Computer_Science/ComputerSystemsA/CSA-Dist2/secretstrings/wordlist")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		request := stubs.Request{Message: scanner.Text()}
		clients[0].Call(stubs.PremiumReverseHandler, request, response)
		fmt.Println("Responded: " + response.Message)
	}
}

//func runServerCommand(client *rpc.Client, request stubs.Request, response stubs.Response) {
//
//}
func main() {
	server := flag.String("server", "127.0.0.1:8030", "IP:port string to connect to as server")
	flag.Parse()
	servers := strings.Split(*server, ",")
	var clients []*rpc.Client
	for serve := range servers {
		fmt.Println("Server: ", servers[serve])
		client, _ := rpc.Dial("tcp", *server)
		clients = append(clients, client)
	}
	getWords(clients)
	for i := range clients {
		clients[i].Close()
	}

	//TODO: connect to the RPC server and send the request(s)
}
