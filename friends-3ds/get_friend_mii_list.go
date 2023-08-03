// Package protocol implements the Friends 3DS protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	friends_3ds_types "github.com/PretendoNetwork/nex-protocols-go/friends-3ds/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetFriendMiiList sets the GetFriendMiiList handler function
func (protocol *Protocol) GetFriendMiiList(handler func(err error, client *nex.Client, callID uint32, friends []*friends_3ds_types.FriendInfo)) {
	protocol.getFriendMiiListHandler = handler
}

func (protocol *Protocol) handleGetFriendMiiList(packet nex.PacketInterface) {
	if protocol.getFriendMiiListHandler == nil {
		globals.Logger.Warning("Friends3DS::GetFriendMiiList not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	friends, err := parametersStream.ReadListStructure(friends_3ds_types.NewFriendInfo())
	if err != nil {
		go protocol.getFriendMiiListHandler(fmt.Errorf("Failed to read friends from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.getFriendMiiListHandler(nil, client, callID, friends.([]*friends_3ds_types.FriendInfo))
}