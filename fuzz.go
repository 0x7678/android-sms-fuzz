package main

import "fmt"
import "rpc"
import "net"
import "log"
import "http"
import "os"
import "strings"
import "bytes"
import "flag"

type Sms string

type Args struct {
	Sms string
}

type Reply struct {
	BytesWritten int
}


func (t *Sms) Send(args *Args, reply *Reply) os.Error {
	addr, _ := net.ResolveUnixAddr("unix", "/dev/socket/rild-debug")
	conn, _ := net.DialUnix("unix", nil, addr)
	buf := bytes.NewBufferString(args.Sms).Bytes()
	fmt.Printf("the buffer is %d bytes long.\n", uint8(len(buf)))
	conn.Write([]byte{uint8(len(buf)), 0, 0, 0})
	count, _ := conn.Write(buf)
	fmt.Printf("wrote %d bytes.\n", count)
	reply.BytesWritten = count
	return nil
}


func main() {
	fmt.Printf("sending an sms\n")
	nextSms := new(Sms)
	rpc.Register(nextSms)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":9909")
	if e != nil {
		log.Exit("listen error:", e)
	}
	go http.Serve(l, nil)

	client, err := rpc.DialHTTP("tcp", "localhost:9909")
	if err != nil {
		log.Exit("dialing:", err)
	}

	flag.Parse()
	if flag.NArg() == 0 {
		fmt.Printf("USAGE: sendsms PDU\n")
		os.Exit(8)
	}

	input := strings.Join(flag.Args(), "")
	fmt.Printf("PDU: %s\n", input)
	args := &Args{input}
	reply := new(Reply)
	err = client.Call("Sms.Send", args, reply)
	if err != nil {
		log.Exit("arith error:", err)
	}
	fmt.Printf("Sms: %s -> %d\n", args.Sms, reply.BytesWritten)

}
