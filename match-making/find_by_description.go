// Package protocol implements the Match Making protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleFindByDescription(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.FindByDescription == nil {
		globals.Logger.Warning("MatchMaking::FindByDescription not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	strDescription, err := parametersStream.ReadString()
	if err != nil {
		_, errorCode = protocol.FindByDescription(fmt.Errorf("Failed to read strDescription from parameters. %s", err.Error()), packet, callID, "", nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	resultRange, err := nex.StreamReadStructure(parametersStream, nex.NewResultRange())
	if err != nil {
		_, errorCode = protocol.FindByDescription(fmt.Errorf("Failed to read resultRange from parameters. %s", err.Error()), packet, callID, "", nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.FindByDescription(nil, packet, callID, strDescription, resultRange)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
