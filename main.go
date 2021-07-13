// socket client for golang
// https://golangr.com
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"time"
)

func main() {

	// connect to server
	conn, _ := net.Dial("tcp", "127.0.0.1:6341")
	conn2, _ := net.Dial("tcp", "127.0.0.1:6342")
	for {

		// Form green color with length header of 6
		hexCmd := "6" + randomGreen()
		fmt.Printf("Sending green to server 1: %s\n", hexCmd[1:])
		// Send to server 1
		fmt.Fprint(conn, hexCmd)

		// Form red color with length header of 6
		hexCmd = "6" + randomRed()
		fmt.Printf("Sending red to server 2: %s\n", hexCmd[1:])
		// Send to server 2
		fmt.Fprint(conn2, hexCmd)

		// wait for reply from server 1
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message from server 1: " + message)
		// wait for reply from server 2
		message2, _ := bufio.NewReader(conn2).ReadString('\n')
		fmt.Print("Message from server 2: " + message2)
	}

}

func randomGreen() string {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	return "00" + fmt.Sprintf("%02X", r.Intn(255)) + "00"
}

func randomRed() string {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	return fmt.Sprintf("%02X", r.Intn(255)) + "00" + "00"
}
