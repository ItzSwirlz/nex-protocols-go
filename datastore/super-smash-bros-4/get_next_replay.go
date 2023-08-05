// Package protocol implements the Super Smash Bros. 4 DataStore protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetNextReplay sets the GetNextReplay handler function
func (protocol *Protocol) GetNextReplay(handler func(err error, client *nex.Client, callID uint32) uint32) {
	protocol.getNextReplayHandler = handler
}

func (protocol *Protocol) handleGetNextReplay(packet nex.PacketInterface) {
	if protocol.getNextReplayHandler == nil {
		globals.Logger.Warning("DataStoreSuperSmashBros4::GetNextReplay not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.getNextReplayHandler(nil, client, callID)
}
