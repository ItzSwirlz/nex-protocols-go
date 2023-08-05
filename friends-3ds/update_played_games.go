// Package protocol implements the Friends 3DS protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	friends_3ds_types "github.com/PretendoNetwork/nex-protocols-go/friends-3ds/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UpdatePlayedGames sets the UpdatePlayedGames handler function
func (protocol *Protocol) UpdatePlayedGames(handler func(err error, client *nex.Client, callID uint32, playedGames []*friends_3ds_types.PlayedGame) uint32) {
	protocol.updatePlayedGamesHandler = handler
}

func (protocol *Protocol) handleUpdatePlayedGames(packet nex.PacketInterface) {
	if protocol.updatePlayedGamesHandler == nil {
		globals.Logger.Warning("Friends3DS::UpdatePlayedGames not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	playedGames, err := parametersStream.ReadListStructure(friends_3ds_types.NewPlayedGame())
	if err != nil {
		go protocol.updatePlayedGamesHandler(fmt.Errorf("Failed to read playedGames from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.updatePlayedGamesHandler(nil, client, callID, playedGames.([]*friends_3ds_types.PlayedGame))
}
