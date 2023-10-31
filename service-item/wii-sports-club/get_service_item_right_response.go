// Package protocol implements the ServiceItemWiiSportsClub protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetServiceItemRightResponse sets the GetServiceItemRightResponse handler function
func (protocol *Protocol) GetServiceItemRightResponse(handler func(err error, packet nex.PacketInterface, callID uint32, requestID uint32) uint32) {
	protocol.getServiceItemRightResponseHandler = handler
}

func (protocol *Protocol) handleGetServiceItemRightResponse(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getServiceItemRightResponseHandler == nil {
		globals.Logger.Warning("ServiceItemWiiSportsClub::GetServiceItemRightResponse not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	requestID, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.getServiceItemRightResponseHandler(fmt.Errorf("Failed to read requestID from parameters. %s", err.Error()), packet, callID, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.getServiceItemRightResponseHandler(nil, packet, callID, requestID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
