// Package protocol implements the Remote Log Device protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

const (
	// ProtocolID is the protocol ID for the RemoteLogDevice protocol
	ProtocolID = 0x1

	// MethodLog is the method ID for the method Log
	MethodLog = 0x1
)

// Protocol handles the RemoteLogDevice protocol
type Protocol struct {
	Server nex.ServerInterface
	Log    func(err error, packet nex.PacketInterface, callID uint32, strLine string) (*nex.RMCMessage, uint32)
}

// Setup initializes the protocol
func (protocol *Protocol) Setup() {
	protocol.Server.OnData(func(packet nex.PacketInterface) {
		message := packet.RMCMessage()

		if message.IsRequest && message.ProtocolID == ProtocolID {
			protocol.HandlePacket(packet)
		}
	})
}

// HandlePacket sends the packet to the correct RMC method handler
func (protocol *Protocol) HandlePacket(packet nex.PacketInterface) {
	request := packet.RMCMessage()

	switch request.MethodID {
	case MethodLog:
		protocol.handleLog(packet)
	default:
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		fmt.Printf("Unsupported RemoteLogDevice method ID: %#v\n", request.MethodID)
	}
}

// NewProtocol returns a new Remote Log Device protocol
func NewProtocol(server nex.ServerInterface) *Protocol {
	protocol := &Protocol{Server: server}

	protocol.Setup()

	return protocol
}
