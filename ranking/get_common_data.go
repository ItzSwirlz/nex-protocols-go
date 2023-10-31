// Package protocol implements the Ranking protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetCommonData sets the GetCommonData handler function
func (protocol *Protocol) GetCommonData(handler func(err error, packet nex.PacketInterface, callID uint32, uniqueID uint64) uint32) {
	protocol.getCommonDataHandler = handler
}

func (protocol *Protocol) handleGetCommonData(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getCommonDataHandler == nil {
		globals.Logger.Warning("Ranking::GetCommonData not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	uniqueID, err := parametersStream.ReadUInt64LE()
	if err != nil {
		errorCode = protocol.getCommonDataHandler(fmt.Errorf("Failed to read uniqueID from parameters. %s", err.Error()), packet, callID, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.getCommonDataHandler(nil, packet, callID, uniqueID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
