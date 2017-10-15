package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"strings"
)

// Request struct represents an HTTP Request
type Request struct {
	headers []string
	path    string
	method  string
}

var responseCodeMap = map[int]string{
	200: "OK",
	404: "Not Found",
}

func parseRequest(conn net.Conn) (request *Request, err error) {
	reader := bufio.NewReader(conn)
	requestLine, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}

	tokens := strings.Split(requestLine, " ")
	// fmt.Printf("Tokens is %v\n", tokens)

	if len(tokens) != 3 {
		// Handle error
		fmt.Printf("Request line should contain Method Request-URI HTTP-Version")
	}

	headers := parseHeaders(reader)
	return &Request{
		method:  tokens[0],
		path:    tokens[1],
		headers: headers,
	}, nil
}

func parseHeaders(reader *bufio.Reader) []string {
	headers := make([]string, 1)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error reading string from connection")
			return nil
		}
		if len(strings.TrimSpace(line)) == 0 {
			break
		}
		headers = append(headers, line)
	}
	return headers
}

func doRespond(conn net.Conn, req *Request) {
	cwd, err := os.Getwd()

	if err != nil {
		panic("Error trying to get Current working directory")
	}

	filePath := cwd + req.path
	fileContents, err := ioutil.ReadFile(filePath)

	responseCode := 200

	var body string

	if err != nil {
		fmt.Printf("Unable to read file %v", filePath)
		responseCode = 404
		body = "File not found\r\n"
	} else {
		body = string(fileContents) + "\r\n"
	}

	// body = "Hello Daniel!"

	statusLine := fmt.Sprintf("HTTP/1.1 %v %v\r\n", responseCode, responseCodeMap[responseCode])

	fmt.Printf("Status line is: %v\n", statusLine)
	fmt.Printf("Body is: %v\n", string(body))

	fmt.Fprintf(conn, statusLine)
	fmt.Fprintf(conn, "Content-Length: %v\r\n", len(body))
	fmt.Fprintf(conn, "\r\n")
	fmt.Fprintf(conn, body)
	// conn.Write([]byte(statusLine))
	// conn.Write(body)
	// conn.Write([]byte("\r\n\r\n"))
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	req, err := parseRequest(conn)
	if err != nil {
		fmt.Printf("Error parsing Request")
	}

	doRespond(conn, req)

	// fmt.Printf("Request is: %v", req)
	// conn.Write([]byte("Hello World!\n"))
}

func startServer() {
	ln, err := net.Listen("tcp", ":9000")

	defer ln.Close()

	if err != nil {
		fmt.Printf("Error opening socket on port 9000")
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Printf("There was an error accepting the connection")
		}
		go handleConnection(conn)
	}
}

func mainHTTPServer() {
	startServer()
}
