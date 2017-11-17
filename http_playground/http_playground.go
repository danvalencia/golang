package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

// Counter is a counter struct
type Counter int

// Chan is a great channel
type Chan chan *http.Request

func (ctr *Counter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	*ctr++
	fmt.Fprintf(w, "counter = %d\n", *ctr)
}

func (ch Chan) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ch <- req
	fmt.Fprint(w, "notification sent!")
}

// ArgServer is a func
func ArgServer(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "<body><div>%v</div></body>", os.Args)
}

func doSomething() (a, b, c, d string) {
	return "a", "b", "c", "d"
}

// MaxOutstanding is the max amount of requests that our system can support
const MaxOutstanding int = 10

var sem = make(chan int, MaxOutstanding)

// Request is a simple function
type Request func()

func process(r Request) {
	r()
}

func handle(r Request) {
	sem <- 1   // Wait for active queue to drain.
	process(r) // May take a long time.
	<-sem      // Done; enable next request to run.
}

// Serve will start a server
func Serve(queue chan Request) {
	for {
		req := <-queue
		go handle(req) // Don't wait for handle to finish.
	}
}

func main() {
	queue := make(chan Request)
	go Serve(queue)

	for i := 0; i < 100; i++ {
		fmt.Printf("Processing line %v\n", i)
		queue <- func(ii int)(req Request) {
			return func() {
			v := ii
			time.Sleep(1000)
			fmt.Printf("Running func %v\n", v)
		}}(i)
	}
}

// func main() {

// 	_, b, _, d := doSomething()

// 	fmt.Printf("%v,%v\n", b, d)
// ch := make(Chan, 10)

// go func() {
// 	for {
// 		req := <-ch
// 		fmt.Printf("Request is: %v", req)
// 	}
// }()

// http.Handle("/args", http.HandlerFunc(ArgServer))
// log.Fatal(http.ListenAndServe(":8080", nil))
// rw := &bufio.ReadWriter{}

// buffer := make([]byte, 10)
// reader := strings.NewReader("abcdefg")

// s := string([]byte{80})
// fmt.Println(s)

// bufReader := bufio.NewReader(reader)

// n, err := bufReader.Read(buffer)

// if err != nil {
// 	fmt.Println("String read successfully")
// }

// fmt.Printf("Number of bytes read: %v\n", n)

// list := []int{4,5,6,3,2,1,1,-1,2,-9,-8,-199,67,10,40}
// }
