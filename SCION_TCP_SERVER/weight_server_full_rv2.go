// weightserver application
// For documentation on how to setup and run the application see:
// https://github.com/netsec-ethz/scion-apps/blob/master/README.md
package main

import (
	//"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	//"os"
	//"strings"

	"sync"

	"github.com/MichaelS11/go-hx711"

	"github.com/scionproto/scion/go/lib/sciond"
	"github.com/scionproto/scion/go/lib/snet"
)

type Rfidth struct {
	Humidity    int
	UID         string
	Temperature int
}

type Message struct {
	Name       string
	Exp        string
	Time       string
	Unitweight float64
	Capacity   int
	CurrentNo  int
	TempL      int
	TempH      int
	HmdL       int
	HmdH       int
	Temp       int
	Hmd        int
	Wght       int
	UID        string
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

var weightData map[string]string
var weightDataLock sync.Mutex

func init() {
	weightData = make(map[string]string)
}

// Obtains input from weight observation application

func printUsage() {
	fmt.Println("weightserver -s ServerSCIONAddress")
	fmt.Println("The SCION address is specified as ISD-AS,[IP Address]:Port")
	fmt.Println("Example SCION address 17-ffaa:0:1102,[192.33.93.173]:42002")
}

func main() {

	var (
		serverAddress  string
		sciondPath     string
		sciondFromIA   bool
		dispatcherPath string

		err    error
		server *snet.Addr

		udpConnection snet.Conn
	)

	var msg = []Message{Message{"Sallos", "10/10/2020", "", 5.71, 21, 0, 0, 40, 10, 90, 20, 30, 0, "[249, 20, 56, 86, 131]"},
		Message{"Toffee", "20/10/2020", "", 4, 50, 0, 0, 50, 10, 90, 20, 30, 0, "[41, 106, 118, 72, 125]"}}
	var noPallette = Message{"NoPallete", "", "", 0, 1, 0, 0, 50, 0, 100, 0, 0, 0, "[]"}
	var sendData Message
	sendData = noPallette
	var rfth = Rfidth{UID: "[]"}
	// Fetch arguments from command line
	flag.StringVar(&serverAddress, "s", "", "Server SCION Address")
	flag.StringVar(&sciondPath, "sciond", "", "Path to sciond socket")
	flag.BoolVar(&sciondFromIA, "sciondFromIA", false, "SCIOND socket path from IA address:ISD-AS")
	flag.StringVar(&dispatcherPath, "dispatcher", "/run/shm/dispatcher/default.sock",
		"Path to dispatcher socket")
	flag.Parse()

	// Create the SCION UDP socket
	if len(serverAddress) > 0 {
		server, err = snet.AddrFromString(serverAddress)
		check(err)
	} else {
		printUsage()
		check(fmt.Errorf("Error, server address needs to be specified with -s"))
	}

	if sciondFromIA {
		if sciondPath != "" {
			log.Fatal("Only one of -sciond or -sciondFromIA can be specified")
		}
		sciondPath = sciond.GetDefaultSCIONDPath(&server.IA)
	} else if sciondPath == "" {
		sciondPath = sciond.GetDefaultSCIONDPath(nil)
	}
	snet.Init(server.IA, sciondPath, dispatcherPath)
	udpConnection, err = snet.ListenSCION("udp4", server)
	check(err)

	defer udpConnection.Close()

	err = hx711.HostInit()

	hx711, err := hx711.NewHx711("GPIO6", "GPIO5")
	check(err)
	// SetGain default is 128
	// Gain of 128 or 64 is input channel A, gain of 32 is input channel B
	// hx711.SetGain(128)
	// make sure to use your values from calibration above
	hx711.AdjustZero = -67280
	hx711.AdjustScale = -2020

	tcpreq := "hello from client"

	tcpAddr, err := net.ResolveTCPAddr("tcp", "0.0.0.0:10000")
	if err != nil {
		panic(err)
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	defer conn.Close()
	if err != nil {
		panic(err)
	}

	receivePacketBuffer := make([]byte, 2500)
	sendPacketBuffer := make([]byte, 3000)
	tcpReply := make([]byte, 1024)

	for {
		_, clientAddress, err := udpConnection.ReadFrom(receivePacketBuffer)
		check(err)

		// Packet received, send back response to same client

		//time.Sleep(200 * time.Microsecond)
		conn.Write([]byte(tcpreq))

		n, err := conn.Read(tcpReply)
		check(err)
		err = json.Unmarshal(tcpReply[:n], &rfth)
		check(err)
		if rfth.UID != "[]" {
			for i := 0; i < len(msg); i++ {
				if string(rfth.UID) == (msg[i].UID) {
					sendData = msg[i]
					break
				}
			}
		}
		data, err := hx711.ReadDataMedian(11)
		check(err)

		if string(sendData.UID) != "[]" {
			sendData.CurrentNo = int(data / sendData.Unitweight)
			sendData.Time = string(time.Now().Format("Jan 2 15:04:05"))
		} else {
			sendData = noPallette
		}
		sendData.Temp = rfth.Temperature
		sendData.Hmd = rfth.Humidity
		sendData.Wght = int(data)
		b, _ := json.Marshal(sendData)
		fmt.Println(string(b))

		copy(sendPacketBuffer, b)

		for i := 0; i < 3; i++ {
			_, err = udpConnection.WriteTo(sendPacketBuffer[:len(b)], clientAddress)
			if err == nil {
				break
			} else {
				log.Println(err)
			}
		}
		check(err)
		sendData = noPallette
	}

}
