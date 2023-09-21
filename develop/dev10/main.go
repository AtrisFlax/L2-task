package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	host, port, err := getHostPort()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	timeout := flag.String("timeout", "10s", "timeout in seconds")
	flag.Parse()

	timoutDuration := getDuration(*timeout)

	address := host + ":" + port
	connection, err := connect(address, timoutDuration)
	defer func(connection net.Conn) {
		_ = connection.Close()
	}(connection)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	//connection -> output
	go copyStream(os.Stdout, connection)
	//input -> connection
	//blocking func. If EOF (ctrl+d) main routine is ended
	copyStream(connection, os.Stdin)
}

func connect(address string, timeout time.Duration) (net.Conn, error) {
	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func copyStream(writer io.Writer, reader io.Reader) {
	if _, err := io.Copy(writer, reader); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func getHostPort() (host string, port string, err error) {
	switch len(os.Args) {
	case 3:
		host = os.Args[2]
		port = "44"
		err = nil
	case 4:
		host = os.Args[2]
		port = os.Args[3]
		err = nil
	default:
		host = ""
		port = ""
		err = fmt.Errorf("wrong invocation format: telnet --timeout=10s [host [port]]")
	}
	return
}

func getDuration(input string) time.Duration {
	timeUnitMap := map[string]time.Duration{
		"d": time.Hour * 24,
		"h": time.Hour,
		"m": time.Minute,
		"s": time.Second,
	}

	timeUnitParts := strings.Split(input, "")

	totalDuration := time.Duration(0)

	var currentNumber string
	var currentUnit string
	for _, part := range timeUnitParts {
		if isNumber(part) {
			currentNumber += part
		} else {
			currentUnit = part
			unitDuration := timeUnitMap[currentUnit]
			number, _ := strconv.Atoi(currentNumber)
			duration := time.Duration(number) * unitDuration
			totalDuration += duration

			currentNumber = ""
			currentUnit = ""
		}
	}
	return totalDuration
}

func isNumber(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}
