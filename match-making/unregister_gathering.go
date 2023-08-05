// Package protocol implements the Match Making protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UnregisterGathering sets the UnregisterGathering handler function
func (protocol *Protocol) UnregisterGathering(handler func(err error, client *nex.Client, callID uint32, idGathering uint32) uint32) {
	protocol.unregisterGatheringHandler = handler
}

func (protocol *Protocol) handleUnregisterGathering(packet nex.PacketInterface) {
	if protocol.unregisterGatheringHandler == nil {
		globals.Logger.Warning("MatchMaking::UnregisterGathering not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	idGathering, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.unregisterGatheringHandler(fmt.Errorf("Failed to read gatheringID from parameters. %s", err.Error()), client, callID, 0)
	}

	go protocol.unregisterGatheringHandler(nil, client, callID, idGathering)
}
