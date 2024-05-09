package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"strconv"

	"github.com/mpschorr/mcp/internal/datatype"
)

func main() {

	bufio.NewWriter(nil)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting TCP server on port 8080")

	for {
		conn, err := l.Accept()
		if err != nil {
			panic(err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	var protocolState int32

	for {
		fmt.Println("- Reading packet")

		length, err1 := datatype.ReadVarInt(reader)
		if errors.Is(err1, io.EOF) {
			fmt.Println("Socket closed by client")
			break
		}
		
		packet, _ := datatype.ReadVarInt(reader)

		fmt.Println("Packet length", length)
		fmt.Println("Packet ID", "0x" + strconv.FormatInt(int64(packet), 16))

		switch packet {

			case 0:
				switch protocolState {
					// Handshake
					case 0:
						protocol, _ := datatype.ReadVarInt(reader)
						address := datatype.ReadString(reader)
						port := datatype.ReadUnsignedShort(reader)
						fmt.Printf("! Ping from %s:%d (protocol version %d)\n", address, port, protocol)
						protocolState, _ = datatype.ReadVarInt(reader)
					
						// Status
					case 1:
						response := createStatusResponse(1, int32(^uint32(0) >> 1), ":3")
						b, _ := json.Marshal(response)

						datatype.WriteVarInt(writer, int32(len(b) + 3)) 
						datatype.WriteVarInt(writer, 0x00)
						datatype.WriteString(writer, string(b))
						writer.Flush()
						protocolState = 0
				}

			// Ping packet
			case 1:
				code := datatype.ReadLong(reader)
				datatype.WriteVarInt(writer, 9) 
				datatype.WriteVarInt(writer, 0x01)
				datatype.WriteLong(writer, code)
				writer.Flush()
		}
	}
}

func createStatusResponse(online int32, max int32, text string) statusResponse {
	return statusResponse{
		Version: statusResponseVersion{
			Name: "1.19.4",
			Protocol: 765,
		},
		Players: &statusResponsePlayers{
			Max: max,
			Online: online,
			Sample: []statusResponsePlayerSample{
				{
					Id: "5617654f-b70d-4e7d-8f18-631e0366bd29",
					Name: "jeelzzz",
				},
			},
		},
		Description: &statusResponseDescription{
			Text: text,
		},
		EnforcesSecureChat: false,
		PreviewsChat: false,
	}
}

type statusResponse struct {
	Version statusResponseVersion `json:"version"`
	Players *statusResponsePlayers `json:"players"`
	Description *statusResponseDescription `json:"description"`
	Favicon string `json:"favicon,omitempty"`
	EnforcesSecureChat bool `json:"enforcesSecureChat"`
	PreviewsChat bool `json:"previewsChat"`
}

type statusResponseVersion struct {
	Name string `json:"name"`
	Protocol int32 `json:"protocol"`
}

type statusResponsePlayers struct {
	Max int32 `json:"max"`
	Online int32 `json:"online"`
	Sample []statusResponsePlayerSample `json:"sample,omitempty"`
}

type statusResponsePlayerSample struct {
	Id string `json:"id"`
	Name string `json:"name"`
}

// TODO should be text component 
type statusResponseDescription struct {
	Text string `json:"text"`
}