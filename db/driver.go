package db

import (
	"bufio"
//	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/url"

	"github.com/aleksysto/golang-db-connect.git/packer"
)

//	type BoltAgent struct {
//		Product string `json:"product"`
//	}
type helloData struct {
	UserData  AuthData `json:"auth"`
	UserAgent string   `json:"user_agent"`
	//BoltAgent BoltAgent `json:"bolt_agent"`
}
type AuthData struct {
	Scheme    string `json:"scheme"`
	Principal string `json:"principal"`
	Password  string `json:"credentials"`
}

const (
	msgReset      byte = 0x0f
	msgRun        byte = 0x10
	msgDiscardAll byte = 0x2f
	msgDiscardN        = msgDiscardAll // Different name >= 4.0
	msgPullAll    byte = 0x3f
	msgPullN           = msgPullAll // Different name >= 4.0
	msgRecord     byte = 0x71
	msgSuccess    byte = 0x70
	msgIgnored    byte = 0x7e
	msgFailure    byte = 0x7f
	msgHello      byte = 0x01
	msgLogon      byte = 0x6A
	msgLogoff     byte = 0x6B
	msgGoodbye    byte = 0x02
	msgBegin      byte = 0x11
	msgCommit     byte = 0x12
	msgRollback   byte = 0x13
	msgRoute      byte = 0x66 // > 4.2
	msgTelemetry  byte = 0x54
)

var (
	magicPreamble = []byte{0x60, 0x60, 0xb0, 0x17}
	versions      = []byte{
		0x00, 0x00, 0x00, 0x05,
		0x00, 0x02, 0x03, 0x04,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
	}
	handshake = append(magicPreamble, versions...)
)

type Driver struct {
	DbUser *string
	DbPass *string
	DbUri  *string
    Decoder packer.Decoder
}

func (d *Driver) Connect(uri string) {
	d.DbUri = &uri
}

func (d *Driver) Execute(query string) {
	fmt.Println(query)
}
func (d *Driver) OpenConnection(uri string) {
	decodedUri, err := url.Parse(uri)
	if err != nil {
		fmt.Println(err)
	}

	host := decodedUri.Host
	fmt.Println(host)
	conn, err := net.Dial("tcp", host)
	if err != nil {
		fmt.Println("error netdial", err)
	}
	//    response := make([]byte, 94) // Assuming a fixed 4-byte response for the handshake
	reader := bufio.NewReader(conn)

	_, err = conn.Write(handshake)
	if err != nil {
		fmt.Println("error handshake write", err)
	}

    res := make([]byte, 2)
    _, _ = reader.Read(res)
    length, err := packer.ReadLength(res)
    fmt.Println(length)
    if err != nil {
        fmt.Println("error scanning bytes")
    }

	response := make([]byte, length)
    n, err := reader.Read(response)
    if err != nil {
        fmt.Println("error reading chunk")
    }
    if uint16(n) != length {
        fmt.Println("error reading all of data")
    }
    fmt.Println(response)
	helloMessage := helloData{
		AuthData{
			"basic",
			"neo4j",
			"password",
		},
		"Go Driver/5.17.0",
		// BoltAgent{},
	}
	fmt.Println(helloMessage)
	helloMessage2 := map[string]any{}
	helloMessage2["user_agent"] = "Go Driver/5.17.0"
	helloMessage2["auth"] = map[string]string{
		"scheme":      "basic",
		"principal":   "neo4j",
		"credentials": "password",
	}
	jsonBytes, err := json.Marshal(helloMessage2)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	helloSlice := append([]byte{0x00, msgHello}, jsonBytes...)

	_, err = conn.Write(helloSlice)
	if err != nil {
		fmt.Println("Error sending hello", err)
		return
	}
    res = make([]byte, 2)
    _, _ = reader.Read(res)
    fmt.Printf("%x", res)
    length, err = packer.ReadLength(res)
    fmt.Println(length)
    if err != nil {
        fmt.Println("error scanning bytes")
    }

	response = make([]byte, length)
    n, err = reader.Read(response)
    if err != nil {
        fmt.Println("error reading chunk")
    }
    if uint16(n) != length {
        fmt.Println("error reading all of data")
    }
    fmt.Printf("%x", response)
	packer.Unmarshal(response)
	conn.Close()
}

func NewDriver() *Driver {
	return &Driver{}
}
