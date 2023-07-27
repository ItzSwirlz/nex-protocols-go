// Package protocol implements the Secure Connection protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

const (
	// ProtocolID is the protocol ID for the Secure Connection protocol
	ProtocolID = 0xB

	// MethodRegister is the method ID for the method Register
	MethodRegister = 0x1

	// MethodRequestConnectionData is the method ID for the method RequestConnectionData
	MethodRequestConnectionData = 0x2

	// MethodRequestURLs is the method ID for the method RequestURLs
	MethodRequestURLs = 0x3

	// MethodRegisterEx is the method ID for the method RegisterEx
	MethodRegisterEx = 0x4

	// MethodTestConnectivity is the method ID for the method TestConnectivity
	MethodTestConnectivity = 0x5

	// MethodUpdateURLs is the method ID for the method UpdateURLs
	MethodUpdateURLs = 0x6

	// MethodReplaceURL is the method ID for the method ReplaceURL
	MethodReplaceURL = 0x7

	// MethodSendReport is the method ID for the method SendReport
	MethodSendReport = 0x8
)

// Protocol stores all the RMC method handlers for the Secure Connection protocol and listens for requests
type Protocol struct {
	Server                       *nex.Server
	registerHandler              func(err error, client *nex.Client, callID uint32, vecMyURLs []*nex.StationURL)
	requestConnectionDataHandler func(err error, client *nex.Client, callID uint32, cidTarget uint32, pidTarget uint32)
	requestURLsHandler           func(err error, client *nex.Client, callID uint32, cidTarget uint32, pidTarget uint32)
	registerExHandler            func(err error, client *nex.Client, callID uint32, vecMyURLs []*nex.StationURL, hCustomData *nex.DataHolder)
	testConnectivityHandler      func(err error, client *nex.Client, callID uint32)
	updateURLsHandler            func(err error, client *nex.Client, callID uint32, vecMyURLs []*nex.StationURL)
	replaceURLHandler            func(err error, client *nex.Client, callID uint32, target *nex.StationURL, url *nex.StationURL)
	sendReportHandler            func(err error, client *nex.Client, callID uint32, reportID uint32, reportData []byte)
}

// Setup initializes the protocol
func (protocol *Protocol) Setup() {
	protocol.Server.On("Data", func(packet nex.PacketInterface) {
		request := packet.RMCRequest()

		if request.ProtocolID() == ProtocolID {
			protocol.HandlePacket(packet)
		}
	})
}

// HandlePacket sends the packet to the correct RMC method handler
func (protocol *Protocol) HandlePacket(packet nex.PacketInterface) {
	request := packet.RMCRequest()

	switch request.MethodID() {
	case MethodRegister:
		go protocol.handleRegister(packet)
	case MethodRequestConnectionData:
		go protocol.handleRequestConnectionData(packet)
	case MethodRequestURLs:
		go protocol.handleRequestURLs(packet)
	case MethodRegisterEx:
		go protocol.handleRegisterEx(packet)
	case MethodTestConnectivity:
		go protocol.handleTestConnectivity(packet)
	case MethodUpdateURLs:
		go protocol.handleUpdateURLs(packet)
	case MethodReplaceURL:
		go protocol.handleReplaceURL(packet)
	case MethodSendReport:
		go protocol.handleSendReport(packet)
	default:
		go globals.RespondNotImplemented(packet, ProtocolID)
		fmt.Printf("Unsupported SecureConnection method ID: %#v\n", request.MethodID())
	}
}

// NewProtocol returns a new Secure Connection protocol
func NewProtocol(server *nex.Server) *Protocol {
	protocol := &Protocol{Server: server}

	protocol.Setup()

	return protocol
}
