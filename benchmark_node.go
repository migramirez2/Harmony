package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"./consensus"
	"./p2p"
)

// Consts
const (
	Message       = "Ping"
	StopCharacter = "\r\n\r\n"
)

// Start a server and process the request by a handler.
func startServer(port int, handler func(net.Conn, consensus.Consensus), consensus consensus.Consensus) {
	listen, err := net.Listen("tcp4", ":"+strconv.Itoa(port))
	defer listen.Close()
	if err != nil {
		log.Fatalf("Socket listen port %d failed,%s", port, err)
		os.Exit(1)
	}
	log.Printf("Begin listen port: %d", port)
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatalln(err)
			continue
		}
		go handler(conn, consensus)
	}
}

func parseMessageType(receivedMessage string) consensus.MessageType {
	return 4
}

func relayToPorts(msg string, conn net.Conn) {
	w := bufio.NewWriter(conn)
	ports := convertIntoInts(msg)
	ch := make(chan int)
	for i, port := range ports {
		go Send(port, strconv.Itoa(i), ch)
	}
	count := 0
	for count < len(ports) {
		fmt.Println(<-ch)
		count++
	}
	w.Write([]byte(Message))
	w.Flush()
	log.Printf("Send: %s", Message)

}

// Helper library to convert '1,2,3,4' into []int{1,2,3,4}.
func convertIntoInts(data string) []int {
	var res = []int{}
	items := strings.Split(data, " ")
	for _, value := range items {
		intValue, err := strconv.Atoi(value)
		checkError(err)
		res = append(res, intValue)
	}
	return res
}

// Do check error.
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

// SocketClient is to connect a socket given a port and send the given message.
func SocketClient(ip, message string, port int) (res string) {
	addr := strings.Join([]string{ip, strconv.Itoa(port)}, ":")
	conn, err := net.Dial("tcp", addr)

	defer conn.Close()

	if err != nil {
		log.Fatalln(err)
	}

	conn.Write([]byte(message))
	conn.Write([]byte(StopCharacter))
	log.Printf("Send: %s", Message)

	buff := make([]byte, 1024)
	n, _ := conn.Read(buff)
	res = string(buff[:n])
	return
}

//https://gist.github.com/kenshinx/5796276
// Send a message to another node with given port.
func Send(port int, message string, ch chan int) (returnMessage string) {
	ip := "127.0.0.1"
	returnMessage = SocketClient(ip, message, port)
	ch <- port
	fmt.Println(returnMessage)
	return
}

// Handler of the leader node.
func NodeHandler(conn net.Conn, consensus consensus.Consensus) {
	defer conn.Close()
	var (
		buf = make([]byte, 1024)
		r   = bufio.NewReader(conn)
	)

	receivedMessage := ""
ILOOP:
	for {
		n, err := r.Read(buf)
		data := string(buf[:n])
		receivedMessage += data
		switch err {
		case io.EOF:
			break ILOOP
		case nil:
			if consensus.IsLeader {
				log.Println("Leader Node is",consensus.Leader)
				log.Println("[Leader] Received:", data)
			if isTransportOver(data) {
				break ILOOP
				}
			} else {
				log.Println("[Slave] Received:", data)
				break ILOOP
			}
		default:
			if consensus.IsLeader {
				log.Fatalf("[Leader] Receive data failed:%s", err)
			} else {
				log.Fatalf("[Slave] Receive data failed:%s", err)
			}
			return
		}
	}
	receivedMessage = strings.TrimSpace(receivedMessage)
	if consensus.IsLeader {
		consensus.ProcessMessageLeader(parseMessageType(receivedMessage), receivedMessage)
	} else {
		log.Printf("[Slave] Send: %s", receivedMessage)
		consensus.ProcessMessageValidator(0, receivedMessage)
	}
	//relayToPorts(receivedMessage, conn)
}

func isTransportOver(data string) (over bool) {
	over = strings.HasSuffix(data, "\r\n\r\n")
	return
}
func createConsensus(i int, port, Ip string) consensus.Consensus{
	// The first Ip, port passed will be leader.
	var leaderPeer p2p.Peer
	consensus := consensus.Consensus{}
	peer :=  p2p.Peer{Port: port, Ip:Ip}
	if i == 0 {
		consensus.IsLeader = true
		consensus.Leader = peer
		leaderPeer =  peer
	} else {
		consensus.IsLeader = false
		consensus.Leader = leaderPeer
	}
	return consensus
}

func main() {
	
	// This should be read from a file
	Ip := "127.0.0.1"
	portInt := 3000
	numOfNodes := 10

	for i := 0; i < numOfNodes; i++ {
		port := portInt + i 
		portString := strconv.Itoa(port)
		consensus := createConsensus(i,portString,Ip)
		fmt.Println(port)
		fmt.Println(consensus)
		//startServer(port,NodeHandler,consensus)
	}
}
