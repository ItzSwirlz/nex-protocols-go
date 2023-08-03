// Package protocol implements the Friends 3DS protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetFriendPicture sets the GetFriendPicture handler function
func (protocol *Protocol) GetFriendPicture(handler func(err error, client *nex.Client, callID uint32, unknown []uint32)) {
	protocol.getFriendPictureHandler = handler
}

func (protocol *Protocol) handleGetFriendPicture(packet nex.PacketInterface) {
	if protocol.getFriendPictureHandler == nil {
		globals.Logger.Warning("Friends3DS::GetFriendPicture not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	unknown, err := parametersStream.ReadListUInt32LE()
	if err != nil {
		go protocol.getFriendPictureHandler(fmt.Errorf("Failed to read unknown from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.getFriendPictureHandler(nil, client, callID, unknown)
}