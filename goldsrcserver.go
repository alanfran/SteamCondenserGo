package servers

import (
	"fmt"
	"net"
	"time"

	"github.com/alanfran/SteamCondenserGo/helpers"
)

type GoldServer server

// GoldServerResponse implements the A2S_INFO query result.
type GoldServerResponse struct {
	Header      byte
	Protocol    byte
	Hostname    string
	Map         string
	Folder      string
	AppId       int64
	Game        string
	NumPlayers  byte
	MaxPlayers  byte
	Bots        byte
	ServerType  byte
	Environment byte
	Visibility  byte
	Vac         byte
}

// QueryGoldServer takes a server address and returns either a response or an error.
func QueryGoldServer(address string) (Response, error) {
	server := GoldServer{Address: address}

	info, err := server.GetInfo()
	if err != nil {
		return Response{}, err
	}

	response := Response{
		Name:       info.Hostname,
		Map:        info.Map,
		Game:       info.Game,
		Players:    int(info.NumPlayers),
		MaxPlayers: int(info.MaxPlayers),
		Bots:       int(info.Bots),
	}

	return response, err
}

// GetInfo queries a GoldServer and returns either a response or an error.
func (model GoldServer) GetInfo() (GoldServerResponse, error) {
	resp := GoldServerResponse{}

	serverAddr, err := net.ResolveUDPAddr("udp", model.Address)
	if err != nil {
		return resp, err
	}

	socket, err := net.DialUDP("udp", nil, serverAddr)
	if err != nil {
		return resp, err
	}
	defer socket.Close()

	query := helpers.CreateNullTermByteString("TSource Engine Query")
	send := createPacket()
	send = append(send, query...)

	socket.SetDeadline(time.Now().Add(time.Second * 1))

	_, err = socket.Write(send)
	if err != nil {
		return resp, err
	}

	data := make([]byte, 4096)
	_, _, err = socket.ReadFromUDP(data)
	if err != nil {
		return resp, err
	}

	resp.bufferToResponse(data)
	return resp, nil
}

func (resp *GoldServerResponse) bufferToResponse(b []byte) {

	reader := helpers.Init(4, b)
	resp.Header = reader.ReadByte()
	resp.Protocol = reader.ReadByte()
	resp.Hostname = reader.ReadNullTermString()
	resp.Map = reader.ReadNullTermString()
	resp.Folder = reader.ReadNullTermString()
	resp.Game = reader.ReadNullTermString()
	resp.AppId = reader.ReadShort()
	resp.NumPlayers = reader.ReadByte()
	resp.MaxPlayers = reader.ReadByte()
	resp.Bots = reader.ReadByte()
	resp.ServerType = reader.ReadByte()
	resp.Environment = reader.ReadByte()
	resp.Visibility = reader.ReadByte()
	resp.Vac = reader.ReadByte()
}

func createPacket() []byte {
	return []byte("\xFF\xFF\xFF\xFF")
}

// PrintDebug prints the fields of a serverResponse into the console.
func (self GoldServerResponse) PrintDebug() {
	fmt.Println("Header: ", self.Header)
	fmt.Println("Protocol: ", self.Protocol)
	fmt.Println("Hostname: ", self.Hostname)
	fmt.Println("Map: ", self.Map)
	fmt.Println("Folder: ", self.Folder)
	fmt.Println("Game: ", self.Game)
	fmt.Println("AppId: ", self.AppId)
	fmt.Println("Players: ", self.NumPlayers, "/", self.MaxPlayers)
	fmt.Println("Bots: ", self.Bots)
	fmt.Println("Server Type: ", self.ServerType)
	fmt.Println("Vac: ", self.Vac)
}
