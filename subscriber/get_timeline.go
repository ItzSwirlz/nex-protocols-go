// Package protocol implements the Subscriber protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetTimeline sets the GetTimeline handler function
func (protocol *Protocol) GetTimeline(handler func(err error, client *nex.Client, callID uint32, packetPayload []byte)) {
	protocol.getTimelineHandler = handler
}

func (protocol *Protocol) handleGetTimeline(packet nex.PacketInterface) {
	if protocol.getTimelineHandler == nil {
		globals.Logger.Warning("Subscriber::GetTimeline not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	globals.Logger.Warning("Subscriber::GetTimeline STUBBED")

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.getTimelineHandler(nil, client, callID, packet.Payload())
}