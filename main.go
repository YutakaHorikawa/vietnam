package main

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "vietmen"
	app.Usage = "vietmen [command] --port [port number]"
	app.Action = func(c *cli.Context) {
		port := c.String("port")
		startServer(port)
	}

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "port",
			Value: "55555",
			Usage: "port number for redis server",
		},
	}

	app.Run(os.Args)

}

func startServer(port string) {
	service := ":" + port
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)
	listner, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		conn, err := listner.Accept()
		if err != nil {
			continue
		}

		go handleClient(conn)
	}

}

func handleClient(conn net.Conn) {
	defer conn.Close()
	conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	fmt.Println("client accept!")
	messageBuf := make([]byte, 1024)
	messageLen, err := conn.Read(messageBuf)
	checkError(err)

	message := string(messageBuf[:messageLen])

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	conn.Write([]byte("ok"))
	notify(message)
}

func notify(message string) {
	notifyData := strings.Split(message, ",")
	command := "display notification \"" + notifyData[0] + "\" with title " + "\"" + notifyData[1] + "\""
	cmd := exec.Command("osascript", "-e", command)
	cmd.Start()
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "fatal: error: %s", err.Error())
		os.Exit(1)
	}
}
