// Package friends implements the Friends QRV protocol
package friends

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// BlackListByName sets the BlackListByName handler function
func (protocol *FriendsProtocol) BlackListByName(handler func(err error, client *nex.Client, callID uint32, strPlayerName string, uiDetails uint32)) {
	protocol.blackListByNameHandler = handler
}

func (protocol *FriendsProtocol) handleBlackListByName(packet nex.PacketInterface) {
	if protocol.blackListByNameHandler == nil {
		globals.Logger.Warning("Friends::BlackListByName not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	strPlayerName, err := parametersStream.ReadString()
	if err != nil {
		go protocol.blackListByNameHandler(fmt.Errorf("Failed to read strPlayerName from parameters. %s", err.Error()), client, callID, "", 0)
		return
	}

	uiDetails, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.blackListByNameHandler(fmt.Errorf("Failed to read uiDetails from parameters. %s", err.Error()), client, callID, "", 0)
		return
	}

	go protocol.blackListByNameHandler(nil, client, callID, strPlayerName, uiDetails)
}