package main

import "fmt"
import "net"
import "os"
import "strings"
import "bytes"
import "flag"

func main() {
	fmt.Printf("sending an sms\n")
	flag.Parse()
	if flag.NArg() == 0 {
		fmt.Printf("USAGE: sendsms PDU\n")
		os.Exit(8)
	}
	input := strings.Join(flag.Args(), "")
	fmt.Printf("PDU: %s\n", input)

	addr, _ := net.ResolveUnixAddr("unix", "/dev/socket/rild-debug")
	conn, _ := net.DialUnix("unix", nil, addr)
	buf := bytes.NewBufferString(input).Bytes()
	fmt.Printf("the buffer is %d bytes long.\n", uint8(len(buf)))
	conn.Write([]byte{uint8(len(buf)), 0, 0, 0})
	count, _ := conn.Write(buf)
	fmt.Printf("wrote %d bytes.\n", count)
}
