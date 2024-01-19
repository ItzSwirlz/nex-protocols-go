// Package protocol implements the ServiceItemWiiSportsClub protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	service_item_wii_sports_club_types "github.com/PretendoNetwork/nex-protocols-go/service-item/wii-sports-club/types"
)

func (protocol *Protocol) handleEndChallenge(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.EndChallenge == nil {
		globals.Logger.Warning("ServiceItemWiiSportsClub::EndChallenge not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	endChallengeParam := service_item_wii_sports_club_types.NewServiceItemEndChallengeParam()
	err = endChallengeParam.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.EndChallenge(fmt.Errorf("Failed to read endChallengeParam from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.EndChallenge(nil, packet, callID, endChallengeParam)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
