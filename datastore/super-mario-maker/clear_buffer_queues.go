// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	datastore_super_mario_maker_types "github.com/PretendoNetwork/nex-protocols-go/datastore/super-mario-maker/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleClearBufferQueues(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.ClearBufferQueues == nil {
		globals.Logger.Warning("DataStoreSuperMarioMaker::ClearBufferQueues not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	params := types.NewList[*datastore_super_mario_maker_types.BufferQueueParam]()
	params.Type = datastore_super_mario_maker_types.NewBufferQueueParam()
	err = params.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.ClearBufferQueues(fmt.Errorf("Failed to read params from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.ClearBufferQueues(nil, packet, callID, params)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
