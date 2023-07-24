// Package matchmake_referee implements the Matchmake Referee NEX protocol
package matchmake_referee

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetRoundParticipants sets the GetRoundParticipants handler function
func (protocol *MatchmakeRefereeProtocol) GetRoundParticipants(handler func(err error, client *nex.Client, callID uint32, roundId uint64)) {
	protocol.GetRoundParticipantsHandler = handler
}

func (protocol *MatchmakeRefereeProtocol) handleGetRoundParticipants(packet nex.PacketInterface) {
	if protocol.GetRoundParticipantsHandler == nil {
		globals.Logger.Warning("MatchmakeReferee::GetRoundParticipants not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	roundId, err := parametersStream.ReadUInt64LE()
	if err != nil {
		go protocol.GetRoundParticipantsHandler(fmt.Errorf("Failed to read roundId from parameters. %s", err.Error()), client, callID, 0)
		return
	}

	go protocol.GetRoundParticipantsHandler(nil, client, callID, roundId)
}
