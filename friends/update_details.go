// Package protocol implements the Friends QRV protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleUpdateDetails(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.UpdateDetails == nil {
		globals.Logger.Warning("Friends::UpdateDetails not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	uiPlayer := types.NewPrimitiveU32(0)
	err = uiPlayer.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.UpdateDetails(fmt.Errorf("Failed to read uiPlayer from parameters. %s", err.Error()), packet, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	uiDetails := types.NewPrimitiveU32(0)
	err = uiDetails.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.UpdateDetails(fmt.Errorf("Failed to read uiDetails from parameters. %s", err.Error()), packet, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.UpdateDetails(nil, packet, callID, uiPlayer, uiDetails)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
