// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// ResetRateCustomRankingCounter sets the ResetRateCustomRankingCounter handler function
func (protocol *Protocol) ResetRateCustomRankingCounter(handler func(err error, client *nex.Client, callID uint32, applicationID uint32) uint32) {
	protocol.resetRateCustomRankingCounterHandler = handler
}

func (protocol *Protocol) handleResetRateCustomRankingCounter(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.resetRateCustomRankingCounterHandler == nil {
		globals.Logger.Warning("DataStoreSuperMarioMaker::ResetRateCustomRankingCounter not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	applicationID, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.resetRateCustomRankingCounterHandler(fmt.Errorf("Failed to read applicationID from parameters. %s", err.Error()), client, callID, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.resetRateCustomRankingCounterHandler(nil, client, callID, applicationID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}